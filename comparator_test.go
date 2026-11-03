package filter

import (
	"errors"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type cmpPart struct {
	Size string
}

type cmpObj struct {
	Name  string
	Size  string
	Parts []cmpPart
}

var errNotANumber = errors.New("not a number")

func compareNumbers(actual, expected string) (int, error) {
	a, err := strconv.Atoi(actual)
	if err != nil {
		return 0, errNotANumber
	}
	e, err := strconv.Atoi(expected)
	if err != nil {
		return 0, errNotANumber
	}
	return a - e, nil
}

func TestComparator(t *testing.T) {
	objs := []cmpObj{
		{Name: "a", Size: "2", Parts: []cmpPart{{Size: "9"}}},
		{Name: "b", Size: "10", Parts: []cmpPart{{Size: "1"}, {Size: "010"}}},
		{Name: "unset"},
	}
	opts := []Option{WithComparator("size", compareNumbers), WithComparator("parts/size", compareNumbers)}

	tests := []struct {
		name  string
		query string
		want  []string
	}{
		{"eq by value", "filter=(eq,size,02)", []string{"a"}},
		{"neq by value", "filter=(neq,size,02)", []string{"b", "unset"}},
		{"gt by value", "filter=(gt,size,9)", []string{"b"}},
		{"gte by value", "filter=(gte,size,10)", []string{"b"}},
		{"lt by value", "filter=(lt,size,9)", []string{"a"}},
		{"lte by value", "filter=(lte,size,2)", []string{"a"}},
		{"nested path through a slice", "filter=(eq,parts/size,10)", []string{"b"}},
		{"other attribute compares as a string", "filter=(eq,name,a)", []string{"a"}},
		{"combined with another expression", "filter=(gt,size,1);(eq,name,b)", []string{"b"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res, err := FilterList(objs, tt.query, opts...)
			require.NoError(t, err)
			names := []string{}
			for _, o := range res {
				names = append(names, o.Name)
			}
			assert.Equal(t, tt.want, names)
		})
	}
}

func TestComparatorMatchesTheWholePath(t *testing.T) {
	objs := []cmpObj{{Name: "a", Size: "2", Parts: []cmpPart{{Size: "02"}}}}
	res, err := FilterList(objs, "filter=(eq,parts/size,2)", WithComparator("size", compareNumbers))
	require.NoError(t, err)
	assert.Empty(t, res)
}

func TestComparatorError(t *testing.T) {
	objs := []cmpObj{{Name: "a", Size: "2"}}
	_, err := FilterList(objs, "filter=(eq,size,abc)", WithComparator("size", compareNumbers))
	assert.ErrorIs(t, err, errNotANumber)
}

func TestNoComparatorComparesStrings(t *testing.T) {
	objs := []cmpObj{{Name: "a", Size: "2"}}
	res, err := FilterList(objs, "filter=(eq,size,02)")
	require.NoError(t, err)
	assert.Empty(t, res)
}
