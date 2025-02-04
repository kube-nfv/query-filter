package filter

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"strings"

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
}

func newFilterListener[T any](unfilteredList []T) *filterListener[T] {
	// Before filtering accept all the objects
	acceptList := make([]bool, len(unfilteredList))
	for i := 0; i < len(unfilteredList); i++ {
		acceptList[i] = true
	}

	return &filterListener[T]{
		unfilteredList: unfilteredList,
		acceptList:     acceptList,
	}
}

func (l *filterListener[T]) ExitSimpleFilterExprOne(ctx *parser.SimpleFilterExprOneContext) {
	for unfObjIdx, unfObj := range l.unfilteredList {
		ac, err := needAcceptOne(unfObj, ctx.OpOne(), ctx.AllAttrName(), ctx.Value())
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
func needAcceptOne[T any](obj T, op parser.IOpOneContext, attrPath []parser.IAttrNameContext, value parser.IValueContext) (bool, error) {
	val := reflect.ValueOf(obj)
	if val.Kind() == reflect.Pointer {
		val = val.Elem()
	}
	attrPathLen := len(attrPath)

	if val.Kind() == reflect.Struct {
		if attrPathLen == 0 {
			return false, fmt.Errorf("field of the struct type can't be the last element")
		}
		nextElem := attrPath[0].GetText()
		// Traverse over fields to find attrPath[0]
		for i := 0; i < val.NumField(); i++ {
			field := val.Field(i)
			fieldType := val.Type().Field(i)
			if exportedFieldNameToLower(fieldType.Name) == nextElem {
				// Found the field that corresponds to the current attrPath.
				if !field.CanInterface() {
					return false, fmt.Errorf("field \"%s\" is unexported", fieldType.Name)
				}
				res, err := needAcceptOne(field.Interface(), op, attrPath[1:], value)
				if err != nil {
					return false, fmt.Errorf("failed to apply filter for the field \"%s\": %w", fieldType.Name, err)
				}
				return res, nil
			}
		}
		return false, fmt.Errorf("struct doesn't contains \"%s\" field", nextElem)
	}

	if val.Kind() == reflect.Slice || val.Kind() == reflect.Array {
		if attrPathLen == 0 {
			return false, fmt.Errorf("field of the slice type can't be the leaf element")
		}
		// If field is slice or array, at least one element should conforms the attrPath filter.
		for i := 0; i < val.Len(); i++ {
			ac, err := needAcceptOne(val.Index(i).Interface(), op, attrPath, value)
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
	binOp, err := opOneToBinaryPredicate[string](op.GetText())
	if err != nil {
		return false, fmt.Errorf("failed to get binary predicate from option \"%s\": %w", op.GetText(), err)
	}
	strVal, err := convertToString(val)
	if err != nil {
		return false, fmt.Errorf("failed to convert leaf field value to the string: %w", err)
	}

	return binOp(strVal, value.GetText()), nil
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
	if val.Kind() == reflect.Pointer {
		val = val.Elem()
	}

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
