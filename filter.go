package filter

import (
	"fmt"
	"strings"

	"github.com/antlr4-go/antlr/v4"
	parser "github.com/kube-nfv/query-filter/antlr/generate"
)

const filterPrefix = "filter="

// Comparator orders an attribute's value against a filter value: negative when actual is smaller,
// zero when equal, positive when larger. An error fails the filter.
type Comparator func(actual, expected string) (int, error)

// Option configures FilterList.
type Option func(comparators map[string]Comparator)

// WithComparator compares the attribute at attrPath ("a/b") with cmp instead of as a string.
// An empty attribute is unset: only neq matches it, and cmp is not called.
func WithComparator(attrPath string, cmp Comparator) Option {
	return func(comparators map[string]Comparator) {
		comparators[attrPath] = cmp
	}
}

func FilterList[T any](in []T, filterQuery string, opts ...Option) ([]T, error) {

	if in == nil || len(in) == 0 || filterQuery == "" {
		return in, nil
	}

	is := antlr.NewInputStream(filterQuery)

	lexer := parser.NewFilterLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	p := parser.NewFilterParser(stream)

	// Errors refer to the expression, without the "filter=" prefix.
	expr := strings.TrimPrefix(filterQuery, filterPrefix)
	errorListener := &customErrorListener{offset: len(filterQuery) - len(expr)}
	p.RemoveErrorListeners()
	p.AddErrorListener(errorListener)

	tree := p.Filter()

	if err := errorListener.Error(); err != nil {
		return nil, fmt.Errorf("failed to parse filter \"%s\": %w", expr, err)
	}

	comparators := make(map[string]Comparator, len(opts))
	for _, opt := range opts {
		opt(comparators)
	}

	listener := newFilterListener(in, comparators)
	antlr.ParseTreeWalkerDefault.Walk(listener, tree)

	return listener.GetFiltered()
}
