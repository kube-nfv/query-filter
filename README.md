## Overview

Query-filter is a golang library that allows Golang object attribute based filtering in runtime based on the
ETSI GS NFS-SOL 013 5.2.2 specification.

Attribute-based filtering allows to reduce the number of objects returned by a query operation. Typically,
attribute-based filtering is applied to a GET request that reads a resource which represents a list of objects (e.g. child
resources). Only those objects that match the filter are returned as part of the resource representation in the message
content of the GET response.

Attribute-based filtering can test a simple (scalar) attribute of the resource representation against a constant value, for
instance for equality, inequality, greater or smaller than, etc. Attribute-based filtering is requested by adding a set of
URI query parameters, the "attribute-based filtering parameters" or "filter" for short, to a resource URI.


### Specification

```
simpleFilterExprOne := <opOne>","<attrName>["/"<attrName>]*","<value>
simpleFilterExprMulti := <opMulti>","<attrName>["/"<attrName>]*","<value>[","<value>]*
simpleFilterExpr := "("<simpleFilterExprOne>")" | "("<simpleFilterExprMulti>")"
filterExpr := <simpleFilterExpr>[";"<simpleFilterExpr>]*
filter := "filter"=<filterExpr>
opOne := "eq" | "neq" | "gt" | "lt" | "gte" | "lte"
opMulti := "in" | "nin" | "cont" | "ncont"
attrName := string
value := string
where:
* zero or more occurrences
[] "" grouping of expressions to be used with *
quotation marks for marking string constants
<> name separator
| separator of alternatives
```
"AttrName" is the name of one attribute in the data type that defines the representation of the resource. The slash ("/")
character in "simpleFilterExprOne" and " simpleFilterExprMulti" allows concatenation of <attrName> entries to filter
by attributes deeper in the hierarchy of a structured document. The special attribute name "@key" refers to the key of a
map.

### Examples

#### 1
Assume we have an list of Golang objects specified by the type:
```
type ExObj struct {
	Id int
	Weight int
	Parts []ExObjPart
}

type ExObjPart struct {
	Id int
	Color string
}
```
with values:
```
[
    {"id":123, "weight":100, "parts":[{"id":1, "color":"red"}, {"id":2, "color":"green"}]},
    {"id":456, "weight":500, "parts":[{"id":3, "color":"green"}, {"id":4, "color":"blue"}]}
]
```
then we can simply filter them in runtime:
```
var objs []*ExObj := GetListOfObjects()
res := filter.FilterList(objs, "filter=(eq,weight,100)")
...
```

That is extremly usefull when a filter comes in the HTTP query params:
```
Request:
GET …/container?filter=(eq,weight,100)
    Response:
    [
        {"id":123, "weight":100, "parts":[{"id":1, "color":"red"}, {"id":2, "color":"green"}]}
    ]
```

### Implementation Notes

The library uses antlr4 generated code to parse the query filter string.
Antlr4 grammars can be found [here](./antlr/grammar/Filter.g4).

For the code generation it used docker image under a ghcr.io/kube-nfv/antlr:v4.13.2

