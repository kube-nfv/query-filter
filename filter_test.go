package filter

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Object and tests from the ETSI GS NFV-SOL 013 5.2.1

type ExObj struct {
	Id     int
	Weight int
	Parts  []ExObjPart
}

type ExObjPart struct {
	Id    int
	Color string
}

func NewSol13Ex1Obj() []*ExObj {
	return []*ExObj{
		{
			Id:     123,
			Weight: 100,
			Parts: []ExObjPart{
				{Id: 1, Color: "red"},
				{Id: 2, Color: "green"},
			},
		},
		{
			Id:     456,
			Weight: 500,
			Parts: []ExObjPart{
				{Id: 3, Color: "green"},
				{Id: 4, Color: "blue"},
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

// Test with hyphenated values
type SoftwareImage struct {
	Name   string
	Status string
}

func TestHyphenatedValues(t *testing.T) {
	images := []*SoftwareImage{
		{Name: "focal-server-cloudimg-amd64", Status: "ready"},
		{Name: "ubuntu-minimal-cloudimg", Status: "pending"},
		{Name: "debian-server-image", Status: "ready"},
	}

	res, err := FilterList(images, "filter=(eq,name,focal-server-cloudimg-amd64)")
	assert.NoError(t, err)
	assert.Equal(t, 1, len(res))
	assert.Equal(t, "focal-server-cloudimg-amd64", res[0].Name)
}

// Test with fmt.Stringer implementation (like k8s Quantity)
type StringerType struct {
	value string
}

func (s StringerType) String() string {
	return s.value
}

type ResourceSpec struct {
	Memory StringerType
	Cpu    int
}

func TestStringerInterface(t *testing.T) {
	resources := []*ResourceSpec{
		{Memory: StringerType{value: "400M"}, Cpu: 4},
		{Memory: StringerType{value: "1Gi"}, Cpu: 8},
		{Memory: StringerType{value: "12Gi"}, Cpu: 16},
	}

	// Filter by memory using Stringer interface
	res, err := FilterList(resources, "filter=(eq,memory,400M)")
	assert.NoError(t, err)
	assert.Equal(t, 1, len(res))
	assert.Equal(t, "400M", res[0].Memory.String())

	// Filter by CPU
	res2, err := FilterList(resources, "filter=(eq,cpu,16)")
	assert.NoError(t, err)
	assert.Equal(t, 1, len(res2))
	assert.Equal(t, "12Gi", res2[0].Memory.String())
}

// Test nested Stringer types (like the user's protobuf case)
type NestedResource struct {
	VirtualMemSize StringerType
}

type VirtualMemory struct {
	VirtualMemSize StringerType
}

type Flavour struct {
	VirtualMemory VirtualMemory
	NumCPU        int
}

func TestNestedStringerInterface(t *testing.T) {
	flavours := []*Flavour{
		{VirtualMemory: VirtualMemory{VirtualMemSize: StringerType{value: "400M"}}, NumCPU: 4},
		{VirtualMemory: VirtualMemory{VirtualMemSize: StringerType{value: "1Gi"}}, NumCPU: 8},
	}

	// Filter by nested memory value
	res, err := FilterList(flavours, "filter=(eq,virtualMemory/virtualMemSize,400M)")
	assert.NoError(t, err)
	assert.Equal(t, 1, len(res))
	assert.Equal(t, "400M", res[0].VirtualMemory.VirtualMemSize.String())
}
