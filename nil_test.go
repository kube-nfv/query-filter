package filter

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type nilInner struct {
	Size string
}

type nilObj struct {
	Name  string
	Inner *nilInner
	Items []*nilInner
	Label *string
}

func TestNilDoesNotMatch(t *testing.T) {
	label := "x"
	objs := []nilObj{
		{Name: "set", Inner: &nilInner{Size: "2"}, Items: []*nilInner{nil, {Size: "2"}}, Label: &label},
		{Name: "nil", Items: []*nilInner{nil}},
	}
	tests := []struct {
		name  string
		query string
	}{
		{"message on the path", "filter=(eq,inner/size,2)"},
		{"neq on the path", "filter=(neq,inner/size,3)"},
		{"message in a slice", "filter=(eq,items/size,2)"},
		{"pointer leaf", "filter=(eq,label,x)"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res, err := FilterList(objs, tt.query)
			require.NoError(t, err)
			require.Len(t, res, 1)
			assert.Equal(t, "set", res[0].Name)
		})
	}
}
