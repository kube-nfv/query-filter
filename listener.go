package filter

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"strings"

	"github.com/antlr4-go/antlr/v4"
	parser "github.com/kube-nfv/query-filter/antlr/generate"
	"golang.org/x/exp/constraints"
)

type filterListener[T any] struct {
	*parser.BaseFilterListener
	errors         []error
	unfilteredList []T
	// Object accept list. List idx identify object in unfilteredList and value identify
	// if object should be accepted.
	acceptList []bool
	// Comparators by attribute path.
	comparators map[string]Comparator
}

func newFilterListener[T any](unfilteredList []T, comparators map[string]Comparator) *filterListener[T] {
	// Before filtering accept all the objects
	acceptList := make([]bool, len(unfilteredList))
	for i := 0; i < len(unfilteredList); i++ {
		acceptList[i] = true
	}

	return &filterListener[T]{
		unfilteredList: unfilteredList,
		acceptList:     acceptList,
		comparators:    comparators,
	}
}

func (l *filterListener[T]) ExitSimpleFilterExprOne(ctx *parser.SimpleFilterExprOneContext) {
	cmp := l.comparators[attrPathText(ctx.AllAttrName())]
	for unfObjIdx, unfObj := range l.unfilteredList {
		ac, err := needAcceptOne(unfObj, ctx.OpOne(), ctx.AllAttrName(), ctx.Value(), cmp)
		if err != nil {
			l.errors = append(l.errors, fmt.Errorf("failed to proceed rule with index \"%d\": %w", ctx.GetRuleIndex(), err))
		}
		// l.acceptList[unfObjIdx] &= ac
		if !ac {
			l.acceptList[unfObjIdx] = false
		}
	}
}

func (l *filterListener[T]) Error() error {
	return errors.Join(l.errors...)
}

func (l *filterListener[T]) GetFiltered() ([]T, error) {
	out := make([]T, 0)
	for objIdx, shouldAccept := range l.acceptList {
		if shouldAccept {
			out = append(out, l.unfilteredList[objIdx])
		}
	}
	return out, l.Error()
}

// Function verifies is the object need to be accepted or filtered
func needAcceptOne[T any](obj T, op parser.IOpOneContext, attrPath []parser.IAttrNameContext, value parser.IValueContext, cmp Comparator) (bool, error) {
	val := derefPointer(reflect.ValueOf(obj))
	attrPathLen := len(attrPath)

	if val.Kind() == reflect.Struct {
		if attrPathLen == 0 {
			if stringer, ok := getStringer(val); ok {
				return applyPredicate(stringer.String(), op.GetText(), valueText(value), cmp)
			}
			return false, fmt.Errorf("field of the struct type can't be the last element")
		}
		nextElem := attrPath[0].GetText()
		hasOneof := false
		for i := 0; i < val.NumField(); i++ {
			field := val.Field(i)
			fieldType := val.Type().Field(i)

			// A protobuf oneof is an interface holding a wrapper struct whose field is the member set.
			if _, ok := fieldType.Tag.Lookup("protobuf_oneof"); ok {
				hasOneof = true
				if !field.IsNil() {
					if wrapper := derefPointer(field.Elem()); wrapper.Kind() == reflect.Struct && wrapper.NumField() == 1 &&
						fieldMatchesName(wrapper.Type().Field(0), nextElem) {
						return needAcceptOne(wrapper.Interface(), op, attrPath, value, cmp)
					}
				}
				continue
			}

			if !fieldMatchesName(fieldType, nextElem) {
				continue
			}

			if !field.CanInterface() {
				return false, fmt.Errorf("field \"%s\" is unexported", fieldType.Name)
			}
			res, err := needAcceptOne(field.Interface(), op, attrPath[1:], value, cmp)
			if err != nil {
				return false, fmt.Errorf("failed to apply filter for the field \"%s\": %w", fieldType.Name, err)
			}
			return res, nil
		}
		if hasOneof {
			// Possibly a oneof member that is not set.
			return false, nil
		}
		return false, fmt.Errorf("struct doesn't contains \"%s\" field", nextElem)
	}

	if val.Kind() == reflect.Slice || val.Kind() == reflect.Array {
		if attrPathLen == 0 {
			return false, fmt.Errorf("field of the slice type can't be the leaf element")
		}
		// If field is slice or array, at least one element should conforms the attrPath filter.
		for i := 0; i < val.Len(); i++ {
			ac, err := needAcceptOne(val.Index(i).Interface(), op, attrPath, value, cmp)
			if err != nil {
				return false, fmt.Errorf("failed to apply filter to the slice element with indedx %d: %w", i, err)
			}
			if ac {
				return true, nil
			}
		}
		// all of the slice ellements doesn't conforms the filter.
		return false, nil
	}

	// We found leaf element which is not struct or list
	if attrPathLen != 0 {
		return false, fmt.Errorf("found leaf field while attributes still exists")
	}
	// Enums (e.g. protobuf) compare by name.
	if stringer, ok := getStringer(val); ok {
		return applyPredicate(stringer.String(), op.GetText(), valueText(value), cmp)
	}
	strVal, err := convertToString(val)
	if err != nil {
		return false, fmt.Errorf("failed to convert leaf field value to the string: %w", err)
	}
	return applyPredicate(strVal, op.GetText(), valueText(value), cmp)
}

type BinaryPredicate[T constraints.Ordered] func(T, T) bool

func opOneToBinaryPredicate[T constraints.Ordered](opOne string) (BinaryPredicate[T], error) {
	switch opOne {
	case "eq":
		return func(t1, t2 T) bool { return t1 == t2 }, nil
	case "neq":
		return func(t1, t2 T) bool { return t1 != t2 }, nil
	case "gt":
		return func(t1, t2 T) bool { return t1 > t2 }, nil
	case "lt":
		return func(t1, t2 T) bool { return t1 < t2 }, nil
	case "gte":
		return func(t1, t2 T) bool { return t1 >= t2 }, nil
	case "lte":
		return func(t1, t2 T) bool { return t1 <= t2 }, nil
	default:
		return nil, fmt.Errorf("undefined opOne operation \"%s\"", opOne)
	}
}

func convertToString(val reflect.Value) (string, error) {
	val = derefPointer(val)
	switch val.Kind() {
	case reflect.String:
		return val.String(), nil
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return strconv.FormatInt(val.Int(), 10), nil
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return strconv.FormatUint(val.Uint(), 10), nil
	case reflect.Float32, reflect.Float64:
		return strconv.FormatFloat(val.Float(), 'f', -1, 64), nil
	case reflect.Bool:
		return strconv.FormatBool(val.Bool()), nil
	default:
		return "", fmt.Errorf("unsupported type: %v", val.Kind())
	}
}

func exportedFieldNameToLower(fieldName string) string {
	if len(fieldName) < 1 {
		return fieldName
	}
	return strings.ToLower(string(fieldName[0])) + fieldName[1:]
}

func derefPointer(val reflect.Value) reflect.Value {
	if val.Kind() == reflect.Pointer {
		return val.Elem()
	}
	return val
}

func getStringer(val reflect.Value) (fmt.Stringer, bool) {
	if stringer, ok := val.Interface().(fmt.Stringer); ok {
		return stringer, true
	}
	if val.CanAddr() {
		if stringer, ok := val.Addr().Interface().(fmt.Stringer); ok {
			return stringer, true
		}
	} else {
		ptrVal := reflect.New(val.Type())
		ptrVal.Elem().Set(val)
		if stringer, ok := ptrVal.Interface().(fmt.Stringer); ok {
			return stringer, true
		}
	}
	return nil, false
}

func applyPredicate(actualValue, op, expectedValue string, cmp Comparator) (bool, error) {
	if cmp != nil {
		return applyComparator(actualValue, op, expectedValue, cmp)
	}
	binOp, err := opOneToBinaryPredicate[string](op)
	if err != nil {
		return false, fmt.Errorf("failed to get binary predicate from option \"%s\": %w", op, err)
	}
	return binOp(actualValue, expectedValue), nil
}

func applyComparator(actualValue, op, expectedValue string, cmp Comparator) (bool, error) {
	binOp, err := opOneToBinaryPredicate[int](op)
	if err != nil {
		return false, fmt.Errorf("failed to get binary predicate from option \"%s\": %w", op, err)
	}
	if actualValue == "" {
		return op == "neq", nil
	}
	res, err := cmp(actualValue, expectedValue)
	if err != nil {
		return false, fmt.Errorf("failed to compare \"%s\" with \"%s\": %w", actualValue, expectedValue, err)
	}
	return binOp(res, 0), nil
}

func attrPathText(attrPath []parser.IAttrNameContext) string {
	names := make([]string, len(attrPath))
	for i, attr := range attrPath {
		names[i] = attr.GetText()
	}
	return strings.Join(names, "/")
}

func fieldMatchesName(field reflect.StructField, name string) bool {
	// Protobuf-generated fields carry their JSON (lowerCamelCase) name in the protobuf tag.
	for _, part := range strings.Split(field.Tag.Get("protobuf"), ",") {
		if jsonName, ok := strings.CutPrefix(part, "json="); ok {
			return jsonName == name
		}
	}
	if jsonTag := field.Tag.Get("json"); jsonTag != "" {
		jsonName := strings.Split(jsonTag, ",")[0]
		if jsonName != "" && jsonName != "-" {
			return jsonName == name
		}
	}
	return exportedFieldNameToLower(field.Name) == name
}

// valueText returns a value as written: unquoted (with ” as a quote), or its original text, spaces included.
func valueText(value parser.IValueContext) string {
	if q := value.QUOTED(); q != nil {
		s := q.GetText()
		return strings.ReplaceAll(s[1:len(s)-1], "''", "'")
	}
	start, stop := value.GetStart(), value.GetStop()
	return start.GetInputStream().GetTextFromInterval(antlr.NewInterval(start.GetStart(), stop.GetStop()))
}
