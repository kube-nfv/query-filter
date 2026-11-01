package filter

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSyntaxErrorReportsTheExpression(t *testing.T) {
	_, err := FilterList([]*Named{{Name: "a"}}, "filter=(eq,name")
	require.Error(t, err)
	assert.NotContains(t, err.Error(), "filter=")
	assert.Contains(t, err.Error(), `"(eq,name"`)
	assert.Contains(t, err.Error(), "line 1:8 ", "column within the expression")
}
