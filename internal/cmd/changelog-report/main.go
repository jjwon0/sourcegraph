package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"flag"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/text"
)

func main() {
	_ = flag.Duration("since", 24*time.Hour, "Report new changelog entries since this period")

	flag.Parse()

	blame, err := parseGitBlame(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}

	changelog, err := parseChangelog(blame)
	if err != nil {
		log.Fatal(err)
	}

	// from := time.Now().UTC().Add(-*since)
	//changelog.Filter(func(c Change) bool {
	//	if c.GitCommit == nil {
	//		log.Fatalf("nil git commit for %+v", c)
	//	}
	//	return !c.GitCommit.AuthorTime.Before(from)
	//})

	json.NewEncoder(os.Stdout).Encode(changelog)
}

type Change struct {
	Description string
	GitCommit   *GitBlameLine
}

type Release struct {
	Version string

	Added   []Change
	Changed []Change
	Fixed   []Change
	Removed []Change
}

type Changelog []Release

func (cl Changelog) Filter(pred func(c Change) bool) {
	for _, r := range cl {
		r.Added = filter(r.Added, pred)
		r.Changed = filter(r.Changed, pred)
		r.Fixed = filter(r.Fixed, pred)
		r.Removed = filter(r.Removed, pred)
	}
}

func filter(cs []Change, pred func(Change) bool) []Change {
	filtered := cs[:0]
	for _, c := range cs {
		if pred(c) {
			filtered = append(filtered, c)
		}
	}
	return filtered
}

func parseChangelog(blame GitBlame) (Changelog, error) {
	p := goldmark.DefaultParser()
	source := blame.Source()
	root := p.Parse(text.NewReader(source))

	var (
		changelog Changelog
		section   *[]Change
		release   *Release
	)

	err := ast.Walk(root, func(n ast.Node, entering bool) (ast.WalkStatus, error) {
		if !entering {
			return ast.WalkContinue, nil
		}

		switch n := n.(type) {
		case *ast.Heading:
			heading := string(bytes.TrimSpace(n.Text(source)))

			switch n.Level {
			case 2:
				if release != nil {
					changelog = append(changelog, *release)
				}
				release = &Release{Version: heading}

				return ast.WalkContinue, nil
			case 3:
				switch heading {
				case "Added":
					section = &release.Added
				case "Changed":
					section = &release.Changed
				case "Fixed":
					section = &release.Fixed
				case "Removed":
					section = &release.Removed
				}

				return ast.WalkSkipChildren, nil
			}

			return ast.WalkContinue, nil

		case *ast.ListItem:
			if section != nil {
				text := n.Text(source)
				*section = append(*section, Change{
					GitCommit:   blame.Find(string(text)),
					Description: string(bytes.TrimSpace(text)),
				})
			}
		}

		return ast.WalkContinue, nil
	})

	return changelog, err
}

type GitBlame []*GitBlameLine

func (b GitBlame) Source() (source []byte) {
	for _, l := range b {
		source = append(source, l.Line...)
		source = append(source, '\n')
	}
	return
}

func (b GitBlame) Find(text string) *GitBlameLine {
	for _, l := range b {
		if strings.Contains(l.Line, text) {
			return l
		}
	}
	return nil
}

type GitBlameLine struct {
	Author     string
	AuthorTime time.Time

	Committer     string
	CommitterTime time.Time

	Ref     string
	Message string

	Line string `json:"-"`
}

// git blame -w --line-porcelain
func parseGitBlame(r io.Reader) (b GitBlame, err error) {
	sc := bufio.NewScanner(r)

	var (
		l = new(GitBlameLine)
		n int
	)

	for sc.Scan() {
		line := sc.Text()
		switch n {
		case 0: // commit ID
			l.Ref = line[:strings.Index(line, " ")]
		case 1:
			l.Author = strings.TrimPrefix(line, "author ")
		case 2:
			l.Author += " " + strings.TrimPrefix(line, "author-mail ")
		case 3:
			ts, _ := strconv.ParseInt(strings.TrimPrefix(line, "author-time "), 10, 64)
			l.AuthorTime = time.Unix(ts, 0).UTC()
		case 4:
			// ignore
		case 5:
			l.Committer = strings.TrimPrefix(line, "committer ")
		case 6:
			l.Committer += " " + strings.TrimPrefix(line, "committer-mail ")
		case 7:
			ts, _ := strconv.ParseInt(strings.TrimPrefix(line, "committer-time "), 10, 64)
			l.CommitterTime = time.Unix(ts, 0).UTC()
		case 8:
			// ignore
		case 9:
			l.Message = strings.TrimPrefix(line, "summary ")
		case 10, 11:
			// ignore
		case 12:
			l.Line = strings.TrimPrefix(line, "\t")
		}

		if n = (n + 1) % 13; n == 0 {
			b = append(b, l)
			l = new(GitBlameLine)
		}
	}

	return b, sc.Err()
}
