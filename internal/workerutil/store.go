package workerutil

import (
	"context"
	"database/sql"
	"time"

	"github.com/keegancsmith/sqlf"
	"github.com/sourcegraph/sourcegraph/internal/db/basestore"
)

// Store is the database layer for the workerutil package that handles worker-side operations.
type Store struct {
	*basestore.Store
	options StoreOptions
}

// StoreOptions configure the behavior of Store over a particular set of tables, columns, and expressions.
type StoreOptions struct {
	// TableName is the name of the table containing work records.
	//
	// The target table (and the target view referenced by `ViewName`) must have the following columns
	// and types:
	//
	//   - id: integer primary key
	//   - state: an enum type containing at least `queued`, `processing`, and `errored`
	//   - failure_message: text
	//   - started_at: timestamp with time zone
	//   - finished_at: timestamp with time zone
	//   - process_after: timestamp with time zone
	//   - num_resets: integer not null
	//
	// It's recommended to put an index or (or partial index) on the state field for more efficient
	// dequeue operations.
	TableName string

	// ViewName is an optional name of a view on top of the table containing work records to query when
	// selecting a candidate and when selecting the record after it has been locked. If this value is
	// not supplied, `TableName` will be used. The value supplied may also indicate a table alias, which
	// can be referenced in `OrderByExpression`, `ColumnExpressions`, and the conditions suplied to
	// `Dequeue`.
	//
	// The target of this field must be a view on top of the configured table with the same column
	// requirements as the base table descried above.
	//
	// Example use case:
	// The processor for LSIF uploads supplies `lsif_uploads_with_repository_name`, a view on top of the
	// `lsif_uploads` table that joins work records with the `repo` table and adds an additional repository
	// name column. This allows `Dequeue` to return a record with additional data so that a second query
	// is not necessary by the caller.
	ViewName string

	// Scan is the function used to convert a rows object into a record of the expected shape.
	Scan RecordScanFn

	// OrderByExpression is the SQL expression used to order candidate records when selecting the next
	// batch of work to perform. This expression may use the alias provided in `ViewName`, if one was
	// supplied.
	OrderByExpression *sqlf.Query

	// ColumnExpressions are the target columns provided to the query when selecting a locked record.
	// These expressions may use the alias provided in `ViewName`, if one was supplied.
	ColumnExpressions []*sqlf.Query

	// StalledMaxAge is the maximum allow duration between updating the state of a record as "processing"
	// and locking the record row during processing. An unlocked row that is marked as processing likely
	// indicates that the worker that dequeued the record has died. There should be a nearly-zero delay
	// between these states during normal operation.
	StalledMaxAge time.Duration

	// MaxNumResets is the maximum number of times a record can be implicitly reset back to the queued
	// state (via `ResetStalled`). If a record's failed attempts counter reaches this threshold, it will
	// be moved into the errored state rather than queued on its next reset to prevent an infinite retry
	// cycle of the same input.
	MaxNumResets int
}

// RecordScanFn is a function that interprets row values as a particular record. This function should
// return a false-valued flag if the given result set was empty. This function must close the rows
// value if the given error value is nil.
//
// See the `CloseRows` function in the store/base package for suggested implementation details.
type RecordScanFn func(rows *sql.Rows, err error) (interface{}, bool, error)

// NewStore creates a new store with the given database handle and options.
func NewStore(handle *basestore.TransactableHandle, options StoreOptions) *Store {
	if options.ViewName == "" {
		options.ViewName = options.TableName
	}

	return &Store{
		Store:   basestore.NewWithHandle(handle),
		options: options,
	}
}

func (s *Store) With(other basestore.ShareableStore) *Store {
	return &Store{Store: s.Store.With(other)}
}

func (s *Store) Transact(ctx context.Context) (*Store, error) {
	txBase, err := s.Store.Transact(ctx)
	if err != nil {
		return nil, err
	}

	return &Store{Store: txBase, options: s.options}, nil
}

// Dequeue selects the first unlocked record matching the given conditions and locks it in a new transaction that
// should be held by the worker process. If there is such an record, it is returned along with a new store instance
// that wraps the transaction. The resulting transaction must be closed by the caller, and the transaction should
// include a state transition of the record into a terminal state. If there is no such unlocked record, a nil record
// and a nil store will be returned along with a  false-valued flag. This method must not be called from within a
// transaction.
//
// The supplied conditions may use the alias provided in `ViewName`, if one was supplied.
func (s *Store) Dequeue(ctx context.Context, conditions []*sqlf.Query) (record interface{}, tx *Store, exists bool, err error) {
	if s.InTransaction() {
		return nil, nil, false, ErrDequeueTransaction
	}

	query := sqlf.Sprintf(
		selectCandidateQuery,
		quote(s.options.ViewName),
		makeConditionSuffix(conditions),
		s.options.OrderByExpression,
		quote(s.options.TableName),
	)

	for {
		// First, we try to select an eligible record outside of a transaction. This will skip
		// any rows that are currently locked inside of a transaction of another dequeue process.
		id, ok, err := basestore.ScanFirstInt(s.Query(ctx, query))
		if err != nil {
			return nil, nil, false, err
		}
		if !ok {
			return nil, nil, false, nil
		}

		// Once we have an eligible identifier, we try to create a transaction and select the
		// record in a way that takes a row lock for the duration of the transaction.
		tx, err = s.Transact(ctx)
		if err != nil {
			return nil, nil, false, err
		}

		// Select the candidate record within the transaction to lock it from other processes. Note
		// that SKIP LOCKED here is necessary, otherwise this query would block on race conditions
		// until the other process has finished with the record.
		_, exists, err = basestore.ScanFirstInt(tx.Query(ctx, sqlf.Sprintf(
			lockQuery,
			quote(s.options.TableName),
			id,
		)))
		if err != nil {
			return nil, nil, false, tx.Done(err)
		}
		if !exists {
			// Due to SKIP LOCKED, This query will return a sql.ErrNoRows error if the record has
			// already been locked in another process's transaction. We'll return a special error
			// that is checked by the caller to try to select a different record.
			if err := tx.Done(ErrDequeueRace); err != ErrDequeueRace {
				return nil, nil, false, err
			}

			// This will occur if we selected a candidate record that raced with another dequeue
			// process. If both dequeue processes select the same record and the other process
			// begins its transaction first, this condition will occur. We'll re-try the process
			// by selecting another identifier - this one will be skipped on a second attempt as
			// it is now locked.
			continue

		}

		// The record is now locked in this transaction. As `TableName` and `ViewName` may have distinct
		// values, we need to perform a second select in order to pass the correct data to the scan
		// function.
		record, exists, err = s.options.Scan(tx.Query(ctx, sqlf.Sprintf(
			selectRecordQuery,
			sqlf.Join(s.options.ColumnExpressions, ", "),
			quote(s.options.ViewName),
			id,
		)))
		if err != nil {
			return nil, nil, false, tx.Done(err)
		}
		if !exists {
			// This only happens on a programming error (mismatch between `TableName` and `ViewName`).
			return nil, nil, false, tx.Done(ErrNoRecord)
		}

		return record, tx, true, nil
	}
}

const selectCandidateQuery = `
-- source: internal/workerutil/store.go:Dequeue
WITH candidate AS (
	SELECT id FROM %s
	WHERE
		state = 'queued' AND
		(process_after IS NULL OR process_after <= NOW())
		%s
	ORDER BY %s
	FOR UPDATE SKIP LOCKED
	LIMIT 1
)
UPDATE %s
SET
	state = 'processing',
	started_at = NOW()
WHERE id IN (SELECT id FROM candidate)
RETURNING id
`

const lockQuery = `
-- source: internal/workerutil/store.go:Dequeue
SELECT 1 FROM %s
WHERE id = %s
FOR UPDATE SKIP LOCKED
LIMIT 1
`

const selectRecordQuery = `
-- source: internal/workerutil/store.go:Dequeue
SELECT %s FROM %s
WHERE id = %s
LIMIT 1
`

// Requeue updates the state of the record with the given identifier to queued and adds a processing delay before
// the next dequeue of this record can be performed.
func (s *Store) Requeue(ctx context.Context, id int, after time.Time) error {
	return s.Exec(ctx, sqlf.Sprintf(
		requeueQuery,
		quote(s.options.TableName),
		after,
		id,
	))
}

const requeueQuery = `
-- source: internal/workerutil/store.go:Requeue
UPDATE %s
SET state = 'queued', process_after = %s
WHERE id = %s
`

// ResetStalled moves all unlocked records in the processing state for more than `StalledMaxAge` back to the queued
// state. In order to prevent input that continually crashes worker instances, records that have been reset more
// than `MaxNumResets` times will be marked as errored. This method returns a list of record identifiers that have
// been reset and a list of record identifiers that have been marked as errored.
func (s *Store) ResetStalled(ctx context.Context) (resetIDs, erroredIDs []int, err error) {
	resetIDs, err = s.resetStalled(ctx, resetStalledQuery)
	if err != nil {
		return nil, nil, err
	}

	erroredIDs, err = s.resetStalled(ctx, resetStalledMaxResetsQuery)
	if err != nil {
		return nil, nil, err
	}

	return resetIDs, erroredIDs, nil
}

func (s *Store) resetStalled(ctx context.Context, q string) ([]int, error) {
	return basestore.ScanInts(s.Query(
		ctx,
		sqlf.Sprintf(
			q,
			quote(s.options.TableName),
			int(s.options.StalledMaxAge/time.Second),
			s.options.MaxNumResets,
			quote(s.options.TableName),
		),
	))
}

const resetStalledQuery = `
-- source: internal/workerutil/store.go:ResetStalled
WITH stalled AS (
	SELECT id FROM %s
	WHERE
		state = 'processing' AND
		NOW() - started_at > (%s * '1 second'::interval) AND
		num_resets < %s
	FOR UPDATE SKIP LOCKED
)
UPDATE %s
SET
	state = 'queued',
	started_at = null,
	num_resets = num_resets + 1
WHERE id IN (SELECT id FROM stalled)
RETURNING id
`

const resetStalledMaxResetsQuery = `
-- source: internal/workerutil/store.go:ResetStalled
WITH stalled AS (
	SELECT id FROM %s
	WHERE
		state = 'processing' AND
		NOW() - started_at > (%s * '1 second'::interval) AND
		num_resets >= %s
	FOR UPDATE SKIP LOCKED
)
UPDATE %s
SET
	state = 'errored',
	finished_at = clock_timestamp(),
	failure_message = 'failed to process'
WHERE id IN (SELECT id FROM stalled)
RETURNING id
`

// quote wraps the given string in a *sqlf.Query so that it is not passed to the database
// as a parameter. It is necessary to quote things such as table names, columns, and other
// expressions that are not simple values.
func quote(s string) *sqlf.Query {
	return sqlf.Sprintf(s)
}

// makeConditionSuffix returns a *sqlf.Query containing "AND {c1 AND c2 AND ...}" when the
// given set of conditions is non-empty, and an empty string otherwise.
func makeConditionSuffix(conditions []*sqlf.Query) *sqlf.Query {
	if len(conditions) == 0 {
		return sqlf.Sprintf("")
	}

	var quotedConditions []*sqlf.Query
	for _, condition := range conditions {
		// Ensure everything is quoted in case the condition has an OR
		quotedConditions = append(quotedConditions, sqlf.Sprintf("(%s)", condition))
	}

	return sqlf.Sprintf("AND %s", sqlf.Join(quotedConditions, " AND "))
}