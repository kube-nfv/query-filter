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
	// offset is subtracted from columns on the first line.
	offset int
}

func (l *customErrorListener) SyntaxError(recognizer antlr.Recognizer, offendingSymbol interface{}, line, column int, msg string, e antlr.RecognitionException) {
	if line == 1 {
		column = max(column-l.offset, 0)
	}
	l.errors = append(l.errors, fmt.Errorf("line %d:%d %s", line, column, msg))
}

func (l *customErrorListener) Error() error {
	return errors.Join(l.errors...)
}
