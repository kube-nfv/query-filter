package filter

import (
	"errors"
	"fmt"

	"github.com/antlr4-go/antlr/v4"
)

var (
	ErrSyntaxError = errors.New("syntax error")
)

type customErrorListener struct {
	*antlr.DefaultErrorListener
	errors []error
}

func (l *customErrorListener) SyntaxError(recognizer antlr.Recognizer, offendingSymbol interface{}, line, column int, msg string, e antlr.RecognitionException) {
	l.errors = append(l.errors, fmt.Errorf("line %d:%d %s", line, column, msg))
}

func (l *customErrorListener) Error() error {
	return errors.Join(l.errors...)
}
