package graphql

import (
	"encoding/json"
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
		for i, x := range v {
			vSlice[i] = x
		}
	case []json.Number:
		vSlice = make([]any, len(v))
		for i, x := range v {
			vSlice[i] = x
		}
	case []bool:
		vSlice = make([]any, len(v))
		for i, x := range v {
			vSlice[i] = x
		}
	case []map[string]any:
		vSlice = make([]any, len(v))
		for i, x := range v {
			vSlice[i] = x
		}
	case []float64:
		vSlice = make([]any, len(v))
		for i, x := range v {
			vSlice[i] = x
		}
	case []float32:
		vSlice = make([]any, len(v))
		for i, x := range v {
			vSlice[i] = x
		}
	case []int:
		vSlice = make([]any, len(v))
		for i, x := range v {
			vSlice[i] = x
		}
	case []int32:
		vSlice = make([]any, len(v))
		for i, x := range v {
			vSlice[i] = x
		}
	case []int64:
		vSlice = make([]any, len(v))
		for i, x := range v {
			vSlice[i] = x
		}
	default:
		vSlice = []any{v}
	}
	return vSlice
}
