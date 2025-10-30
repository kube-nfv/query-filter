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

type SimpleFlavour struct {
	VirtualMemory VirtualMemory
	NumCPU        int
}

func TestNestedStringerInterface(t *testing.T) {
	flavours := []*SimpleFlavour{
		{VirtualMemory: VirtualMemory{VirtualMemSize: StringerType{value: "400M"}}, NumCPU: 4},
		{VirtualMemory: VirtualMemory{VirtualMemSize: StringerType{value: "1Gi"}}, NumCPU: 8},
	}

	// Filter by nested memory value
	res, err := FilterList(flavours, "filter=(eq,virtualMemory/virtualMemSize,400M)")
	assert.NoError(t, err)
	assert.Equal(t, 1, len(res))
	assert.Equal(t, "400M", res[0].VirtualMemory.VirtualMemSize.String())
}

// Test Stringer with pointer receiver (like k8s Quantity)
type PointerStringerType struct {
	value string
}

// String method defined on pointer receiver
func (s *PointerStringerType) String() string {
	return s.value
}

type K8sQuantity struct {
	Size PointerStringerType
}

func TestPointerReceiverStringer(t *testing.T) {
	resources := []*K8sQuantity{
		{Size: PointerStringerType{value: "12Gi"}},
		{Size: PointerStringerType{value: "400M"}},
		{Size: PointerStringerType{value: "1Gi"}},
	}

	// Filter by size using pointer receiver Stringer
	res, err := FilterList(resources, "filter=(eq,size,12Gi)")
	assert.NoError(t, err)
	assert.Equal(t, 1, len(res))
	assert.Equal(t, "12Gi", res[0].Size.String())
}

// Test JSON tag support
type Quantity struct {
	value string
}

func (q *Quantity) String() string {
	return q.value
}

type StorageAttribute struct {
	TypeOfStorage string   `json:"typeOfStorage"`
	SizeOfStorage Quantity `json:"sizeOfStorage"`
	IsBoot        bool     `json:"isBoot"`
}

type VirtualMemorySpec struct {
	VirtualMemSize Quantity `json:"string"`
}

type VirtualCPU struct {
	NumVirtualCpu int `json:"numVirtualCpu"`
}

type Flavour struct {
	VirtualMemory     VirtualMemorySpec  `json:"virtualMemory"`
	VirtualCpu        VirtualCPU         `json:"virtualCpu"`
	StorageAttributes []StorageAttribute `json:"storageAttributes"`
	FlavourId         string             `json:"flavourId"`
}

func TestJSONTagSupport(t *testing.T) {
	flavours := []*Flavour{
		{
			VirtualMemory: VirtualMemorySpec{VirtualMemSize: Quantity{value: "400M"}},
			VirtualCpu:    VirtualCPU{NumVirtualCpu: 4},
			StorageAttributes: []StorageAttribute{
				{TypeOfStorage: "volume", SizeOfStorage: Quantity{value: "12Gi"}, IsBoot: true},
			},
			FlavourId: "flav-1",
		},
		{
			VirtualMemory: VirtualMemorySpec{VirtualMemSize: Quantity{value: "1Gi"}},
			VirtualCpu:    VirtualCPU{NumVirtualCpu: 8},
			StorageAttributes: []StorageAttribute{
				{TypeOfStorage: "volume", SizeOfStorage: Quantity{value: "20Gi"}, IsBoot: false},
			},
			FlavourId: "flav-2",
		},
	}

	// Test filtering by JSON tag name (not Go field name)
	res, err := FilterList(flavours, "filter=(eq,virtualMemory/string,400M)")
	assert.NoError(t, err)
	assert.Equal(t, 1, len(res))
	assert.Equal(t, "flav-1", res[0].FlavourId)

	// Test filtering by nested JSON tag
	res2, err := FilterList(flavours, "filter=(eq,virtualCpu/numVirtualCpu,4)")
	assert.NoError(t, err)
	assert.Equal(t, 1, len(res2))
	assert.Equal(t, "flav-1", res2[0].FlavourId)

	// Test filtering in slice with JSON tag
	res3, err := FilterList(flavours, "filter=(eq,storageAttributes/sizeOfStorage,12Gi)")
	assert.NoError(t, err)
	assert.Equal(t, 1, len(res3))
	assert.Equal(t, "flav-1", res3[0].FlavourId)

	// Test complex filter matching user's scenario
	res4, err := FilterList(flavours, "filter=(eq,virtualMemory/string,400M);(eq,virtualCpu/numVirtualCpu,4);(eq,storageAttributes/sizeOfStorage,12Gi)")
	assert.NoError(t, err)
	assert.Equal(t, 1, len(res4))
	assert.Equal(t, "flav-1", res4[0].FlavourId)
}

// Test that fields without JSON tags still work (backwards compatibility)
type LegacyResource struct {
	Memory int
	Cpu    int
}

func TestBackwardsCompatibility(t *testing.T) {
	resources := []*LegacyResource{
		{Memory: 100, Cpu: 4},
		{Memory: 200, Cpu: 8},
	}

	// Should still work with Go field names
	res, err := FilterList(resources, "filter=(eq,memory,100)")
	assert.NoError(t, err)
	assert.Equal(t, 1, len(res))
	assert.Equal(t, 100, res[0].Memory)
}

// Test JSON tags with fmt.Stringer leaf nodes
type StringerWithJSONTag struct {
	value string
}

func (s *StringerWithJSONTag) String() string {
	return s.value
}

type ResourceWithJSONStringer struct {
	Size StringerWithJSONTag `json:"sizeOfStorage"`
	Name string              `json:"resourceName"`
}

func TestJSONTagWithStringerLeaf(t *testing.T) {
	resources := []*ResourceWithJSONStringer{
		{Size: StringerWithJSONTag{value: "12Gi"}, Name: "resource-1"},
		{Size: StringerWithJSONTag{value: "400M"}, Name: "resource-2"},
	}

	// Filter using JSON tag name on a Stringer field
	res, err := FilterList(resources, "filter=(eq,sizeOfStorage,12Gi)")
	assert.NoError(t, err)
	assert.Equal(t, 1, len(res))
	assert.Equal(t, "resource-1", res[0].Name)
	assert.Equal(t, "12Gi", res[0].Size.String())
}
