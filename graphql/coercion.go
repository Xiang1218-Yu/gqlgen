package graphql

import (
	"encoding/json"
	"reflect"
)

// CoerceList applies coercion from a single value to a list.
func CoerceList(v any) []any {
	var vSlice []any
	if v == nil {
		return vSlice
	}

	switch v := v.(type) {
	case []any:
		// already a slice no coercion required
		vSlice = v
	case []string:
		vSlice = make([]any, len(v))
		for i := range v {
			vSlice[i] = v[i]
		}
	case []json.Number:
		vSlice = make([]any, len(v))
		for i := range v {
			vSlice[i] = v[i]
		}
	case []bool:
		vSlice = make([]any, len(v))
		for i := range v {
			vSlice[i] = v[i]
		}
	case []map[string]any:
		vSlice = make([]any, len(v))
		for i := range v {
			vSlice[i] = v[i]
		}
	case []float64:
		vSlice = make([]any, len(v))
		for i := range v {
			vSlice[i] = v[i]
		}
	case []float32:
		vSlice = make([]any, len(v))
		for i := range v {
			vSlice[i] = v[i]
		}
	case []int:
		vSlice = make([]any, len(v))
		for i := range v {
			vSlice[i] = v[i]
		}
	case []int32:
		vSlice = make([]any, len(v))
		for i := range v {
			vSlice[i] = v[i]
		}
	case []int64:
		vSlice = make([]any, len(v))
		for i := range v {
			vSlice[i] = v[i]
		}
	default:
		// Fall back to reflection so any other slice or array type (e.g.
		// []int8, []uint, []complex64, or a named slice type) keeps every
		// element instead of being treated as a single scalar value.
		rv := reflect.ValueOf(v)
		switch rv.Kind() {
		case reflect.Slice:
			if rv.IsNil() {
				return vSlice
			}
			fallthrough
		case reflect.Array:
			vSlice = make([]any, rv.Len())
			for i := range vSlice {
				vSlice[i] = rv.Index(i).Interface()
			}
		default:
			vSlice = []any{v}
		}
	}
	return vSlice
}
