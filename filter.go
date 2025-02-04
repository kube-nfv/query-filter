package filter

import (
	"fmt"

	"github.com/antlr4-go/antlr/v4"
	parser "github.com/kube-nfv/query-filter/antlr/generate"
)

func FilterList[T any](in []T, filterQuery string) ([]T, error) {

	if in == nil || len(in) == 0 || filterQuery == "" {
		return in, nil
	}

	is := antlr.NewInputStream(filterQuery)

	lexer := parser.NewFilterLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	p := parser.NewFilterParser(stream)

	errorListener := &customErrorListener{}
	p.RemoveErrorListeners()
	p.AddErrorListener(errorListener)

	tree := p.Filter()

	if err := errorListener.Error(); err != nil {
		return nil, fmt.Errorf("failed to parse filter query \"%s\": %w", filterQuery, err)
	}

	listener := newFilterListener(in)
	antlr.ParseTreeWalkerDefault.Walk(listener, tree)

	return listener.GetFiltered()
}
