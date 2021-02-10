package graphqlbackend

import (
	"context"

	"github.com/sourcegraph/sourcegraph/cmd/frontend/graphqlbackend/graphqlutil"
)

type docSymbolsArgs struct {
	LSIFSymbolsArgs
	First *int
}

func (r *GitTreeEntryResolver) DocSymbols(ctx context.Context, args *docSymbolsArgs) (DocSymbolConnectionResolver, error) {
	lsifResolver, err := r.LSIF(ctx, &struct{ ToolName *string }{})
	if err != nil {
		return nil, err
	}
	symbolsConnection, err := lsifResolver.Symbols(ctx, &args.LSIFSymbolsArgs)
	if err != nil {
		return nil, err
	}
	return symbolsConnection, nil
}

type DocSymbolConnectionResolver interface {
	Nodes(ctx context.Context) ([]DocSymbolResolver, error)
	PageInfo(ctx context.Context) (*graphqlutil.PageInfo, error)
}

type DocSymbolResolver interface {
	Text(ctx context.Context) (string, error)
	Detail(ctx context.Context) (string, error)
	Kind(ctx context.Context) (string, error)   /* enum SymbolKind */
	Tags(ctx context.Context) ([]string, error) /* enum SymbolTag */
	Definitions(ctx context.Context) (LocationConnectionResolver, error)
	Children(ctx context.Context) ([]DocSymbolResolver, error)
}