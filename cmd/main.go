package main

import (
	"fmt"

	"github.com/antlr4-go/antlr/v4"
	parser "github.com/kube-nfv/query-filter/antlr/generate"
)

type filterListener struct {
	*parser.BaseFilterListener
}

func main() {
	//in := "filter=(eq,mymap/@key,abc123)"
	in := "filter=(eq,name,John);(lt,age,30)"
	is := antlr.NewInputStream(in)	
	lexer := parser.NewFilterLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	p := parser.NewFilterParser(stream)

	tree := p.Filter()

	fmt.Println(in)
	fmt.Println(tree.ToStringTree(nil, p))
}
