package filter

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Object and tests from the ETSI GS NFV-SOL 013 5.2.1

type ExObj struct {
	Id int
	Weight int	
	Parts []ExObjPart
}

type ExObjPart struct {
	Id int
	Color string
}

func NewSol13Ex1Obj() []*ExObj {
	return []*ExObj{
		{
			Id: 123,
			Weight: 100,
			Parts: []ExObjPart{
				{ Id: 1, Color: "red" },
				{ Id: 2, Color: "green" },
			},
		},
		{
			Id: 456,
			Weight: 500,
			Parts: []ExObjPart{
				{ Id: 3, Color: "green" },
				{ Id: 4, Color: "blue"},
			},
		},
	}
}

func TestSol13Ex2EmptyFilter(t *testing.T) {
	res, err := FilterList(NewSol13Ex1Obj(), "")
	assert.NoError(t, err)
	assert.Equal(t, res, NewSol13Ex1Obj())
}

func TestSol13Ex3SimpleFilter(t *testing.T) {
	res, err := FilterList(NewSol13Ex1Obj(), "filter=(eq,weight,100)")
	assert.NoError(t, err)
	assert.Equal(t, res, []*ExObj{NewSol13Ex1Obj()[0]})
}

func TestSol13Ex4ManyResults(t *testing.T) {
	res, err := FilterList(NewSol13Ex1Obj(), "filter=(eq,parts/color,green)")
	assert.NoError(t, err)
	assert.Equal(t, res, NewSol13Ex1Obj())
}

func TestSol13Ex4MultipleFilters(t *testing.T) {
	res, err := FilterList(NewSol13Ex1Obj(), "filter=(eq,parts/color,green);(eq,parts/id,3)")
	assert.NoError(t, err)
	assert.Equal(t, res, []*ExObj{NewSol13Ex1Obj()[1]})
}
