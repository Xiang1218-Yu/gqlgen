package graphql

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCoerceList(t *testing.T) {
	mapInput := map[string]any{
		"test": "value",
		"nested": map[string]any{
			"nested": true,
		},
	}

	jsonNumber := json.Number("12")

	assert.Equal(t, []any{"test", "values"}, CoerceList([]any{"test", "values"}))
	assert.Equal(t, []any{"test"}, CoerceList("test"))
	assert.Equal(t, []any{"test"}, CoerceList([]string{"test"}))
	// multi-element typed slices must keep every element, not just the first
	assert.Equal(t, []any{"a", "b", "c"}, CoerceList([]string{"a", "b", "c"}))
	assert.Equal(t, []any{1, 2, 3}, CoerceList([]int{1, 2, 3}))
	assert.Equal(t, []any{int32(1), int32(2)}, CoerceList([]int32{1, 2}))
	assert.Equal(t, []any{int64(1), int64(2)}, CoerceList([]int64{1, 2}))
	assert.Equal(t, []any{float32(1.5), float32(2.5)}, CoerceList([]float32{1.5, 2.5}))
	assert.Equal(t, []any{1.5, 2.5}, CoerceList([]float64{1.5, 2.5}))
	assert.Equal(t, []any{jsonNumber, json.Number("34")}, CoerceList([]json.Number{jsonNumber, json.Number("34")}))
	assert.Equal(t, []any{true, false}, CoerceList([]bool{true, false}))
	assert.Equal(t, []any{mapInput, mapInput}, CoerceList([]map[string]any{mapInput, mapInput}))
	// unlisted slice kinds are covered via reflection
	assert.Equal(t, []any{int8(1), int8(2)}, CoerceList([]int8{1, 2}))
	assert.Equal(t, []any{uint(1), uint(2)}, CoerceList([]uint{1, 2}))
	assert.Equal(t, []any{complex64(1), complex64(2)}, CoerceList([]complex64{1, 2}))
	assert.Equal(t, []any{mapInput}, CoerceList([]any{mapInput}))
	assert.Empty(t, CoerceList(nil))
	assert.Empty(t, CoerceList([]string(nil)))
}
