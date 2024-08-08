package kv

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestGet(t *testing.T) {
	var testdata = map[string]any{
		"string":              "hello world",
		"int":                 -1,
		"uint":                uint(1),
		"int8":                int8(-1),
		"uint8":               uint8(1),
		"int16":               int16(-1),
		"uint16":              uint16(1),
		"int32":               int32(-1),
		"uint32":              uint32(1),
		"int64":               int64(-1),
		"uint64":              uint64(1),
		"float32":             float32(1.0),
		"float64":             1.0,
		"bool":                true,
		"duration":            time.Second,
		"time":                time.Date(2024, 7, 18, 0, 0, 0, 0, time.UTC),
		"[]int":               []int{1, 2, 3},
		"[]string":            []string{"a", "b", "c"},
		"[]bool":              []bool{true, false},
		"[]duration":          []time.Duration{time.Second, time.Minute},
		"[]any":               []any{1, "2", 3.0},
		"map[string]any":      map[string]any{"key1": "value1", "key2": 1},
		"map[string]string":   map[string]string{"key": "value"},
		"map[string]bool":     map[string]bool{"key": true, "key2": false},
		"map[string]int":      map[string]int{"key": 1, "key2": 2},
		"map[string]int64":    map[string]int64{"key": 1, "key2": 2},
		"map[string][]string": map[string][]string{"key": {"a", "b", "c"}},
	}
	t.Run("int", func(t *testing.T) {
		assert.Equal(t, -1, Get[int](testdata, "int"))
	})
	t.Run("uint", func(t *testing.T) {
		assert.Equal(t, uint(1), Get[uint](testdata, "uint"))
	})
	t.Run("int8", func(t *testing.T) {
		assert.Equal(t, int8(-1), Get[int8](testdata, "int8"))
	})
	t.Run("uint8", func(t *testing.T) {
		assert.Equal(t, uint8(1), Get[uint8](testdata, "uint8"))
	})
	t.Run("int16", func(t *testing.T) {
		assert.Equal(t, int16(-1), Get[int16](testdata, "int16"))
	})
	t.Run("uint16", func(t *testing.T) {
		assert.Equal(t, uint16(1), Get[uint16](testdata, "uint16"))
	})
	t.Run("int32", func(t *testing.T) {
		assert.Equal(t, int32(-1), Get[int32](testdata, "int32"))
	})
	t.Run("uint32", func(t *testing.T) {
		assert.Equal(t, uint32(1), Get[uint32](testdata, "uint32"))
	})
	t.Run("int64", func(t *testing.T) {
		assert.Equal(t, int64(-1), Get[int64](testdata, "int64"))
	})
	t.Run("uint64", func(t *testing.T) {
		assert.Equal(t, uint64(1), Get[uint64](testdata, "uint64"))
	})
	t.Run("float32", func(t *testing.T) {
		assert.Equal(t, float32(1.0), Get[float32](testdata, "float32"))
	})
	t.Run("float64", func(t *testing.T) {
		assert.Equal(t, float64(1.0), Get[float64](testdata, "float64"))
	})
	t.Run("bool", func(t *testing.T) {
		assert.Equal(t, true, Get[bool](testdata, "bool"))
	})
	t.Run("duration", func(t *testing.T) {
		assert.Equal(t, time.Second, Get[time.Duration](testdata, "duration"))
	})
	t.Run("time", func(t *testing.T) {
		assert.Equal(t, time.Date(2024, 7, 18, 0, 0, 0, 0, time.UTC), Get[time.Time](testdata, "time"))
	})
	t.Run("string", func(t *testing.T) {
		assert.Equal(t, "hello world", Get[string](testdata, "string"))
	})
	t.Run("[]int", func(t *testing.T) {
		assert.Equal(t, []int{1, 2, 3}, Get[[]int](testdata, "[]int"))
	})
	t.Run("[]string", func(t *testing.T) {
		assert.Equal(t, []string{"a", "b", "c"}, Get[[]string](testdata, "[]string"))
	})
	t.Run("[]bool", func(t *testing.T) {
		assert.Equal(t, []bool{true, false}, Get[[]bool](testdata, "[]bool"))
	})
	t.Run("[]duration", func(t *testing.T) {
		assert.Equal(t, []time.Duration{time.Second, time.Minute}, Get[[]time.Duration](testdata, "[]duration"))
	})
	t.Run("[]any", func(t *testing.T) {
		assert.Equal(t, []any{1, "2", 3.0}, Get[[]any](testdata, "[]any"))
	})
	t.Run("map[string]any", func(t *testing.T) {
		assert.Equal(t, map[string]any{"key1": "value1", "key2": 1}, Get[map[string]any](testdata, "map[string]any"))
	})
	t.Run("map[string]string", func(t *testing.T) {
		assert.Equal(t, map[string]string{"key": "value"}, Get[map[string]string](testdata, "map[string]string"))
	})
	t.Run("map[string]bool", func(t *testing.T) {
		assert.Equal(t, map[string]bool{"key": true, "key2": false}, Get[map[string]bool](testdata, "map[string]bool"))
	})
	t.Run("map[string]int", func(t *testing.T) {
		assert.Equal(t, map[string]int{"key": 1, "key2": 2}, Get[map[string]int](testdata, "map[string]int"))
	})
	t.Run("map[string]int64", func(t *testing.T) {
		assert.Equal(t, map[string]int64{"key": 1, "key2": 2}, Get[map[string]int64](testdata, "map[string]int64"))
	})
	t.Run("map[string][]string", func(t *testing.T) {
		assert.Equal(t, map[string][]string{"key": {"a", "b", "c"}}, Get[map[string][]string](testdata, "map[string][]string"))
	})
}

func TestGet_WithConvert(t *testing.T) {
	var testdata = map[string]any{
		"string":     []byte("hello world"),
		"int":        1.0,
		"uint":       1.0,
		"int8":       1.0,
		"uint8":      1.0,
		"int16":      1.0,
		"uint16":     1.0,
		"int32":      1.0,
		"uint32":     1.0,
		"int64":      1.0,
		"uint64":     1.0,
		"float32":    "1",
		"float64":    "1",
		"bool":       1,
		"duration":   1000000000,
		"time":       "2024-08-08",
		"[]int":      []float32{1.0, 2.0, 3.0},
		"[]string":   []int{1, 2, 3},
		"[]bool":     []int{1, 0},
		"[]duration": []int64{1000000000},
		"[]any": []map[string]any{{
			"1": 1,
			"2": 2,
			"3": 3,
		}},
		"map[string]any":      map[any]any{1: "value1", 2: 1},
		"map[string]string":   map[any]any{1: "value"},
		"map[string]bool":     map[any]any{1: true, 2: false},
		"map[string]int":      map[any]any{1: 1, 2: 2},
		"map[string]int64":    map[any]any{1: 1, 2: 2},
		"map[string][]string": map[any][]string{1: {"a", "b", "c"}},
	}
	t.Run("string", func(t *testing.T) {
		assert.Equal(t, "hello world", Get[string](testdata, "string"))
	})
	t.Run("int", func(t *testing.T) {
		assert.Equal(t, int(1), Get[int](testdata, "int"))
	})
	t.Run("uint", func(t *testing.T) {
		assert.Equal(t, uint(1), Get[uint](testdata, "uint"))
	})
	t.Run("int8", func(t *testing.T) {
		assert.Equal(t, int8(1), Get[int8](testdata, "int8"))
	})
	t.Run("uint8", func(t *testing.T) {
		assert.Equal(t, uint8(1), Get[uint8](testdata, "uint8"))
	})
	t.Run("int16", func(t *testing.T) {
		assert.Equal(t, int16(1), Get[int16](testdata, "int16"))
	})
	t.Run("uint16", func(t *testing.T) {
		assert.Equal(t, uint16(1), Get[uint16](testdata, "uint16"))
	})
	t.Run("int32", func(t *testing.T) {
		assert.Equal(t, int32(1), Get[int32](testdata, "int32"))
	})
	t.Run("uint32", func(t *testing.T) {
		assert.Equal(t, uint32(1), Get[uint32](testdata, "uint32"))
	})
	t.Run("int64", func(t *testing.T) {
		assert.Equal(t, int64(1), Get[int64](testdata, "int64"))
	})
	t.Run("uint64", func(t *testing.T) {
		assert.Equal(t, uint64(1), Get[uint64](testdata, "uint64"))
	})
	t.Run("float32", func(t *testing.T) {
		assert.Equal(t, float32(1.0), Get[float32](testdata, "float32"))
	})
	t.Run("float64", func(t *testing.T) {
		assert.Equal(t, 1.0, Get[float64](testdata, "float64"))
	})
	t.Run("bool", func(t *testing.T) {
		assert.Equal(t, true, Get[bool](testdata, "bool"))
	})
	t.Run("duration", func(t *testing.T) {
		assert.Equal(t, time.Second, Get[time.Duration](testdata, "duration"))
	})
	t.Run("time", func(t *testing.T) {
		assert.Equal(t, time.Date(2024, 8, 8, 0, 0, 0, 0, time.UTC), Get[time.Time](testdata, "time"))
	})
	t.Run("[]int", func(t *testing.T) {
		assert.Equal(t, []int{1, 2, 3}, Get[[]int](testdata, "[]int"))
	})
	t.Run("[]string", func(t *testing.T) {
		assert.Equal(t, []string{"1", "2", "3"}, Get[[]string](testdata, "[]string"))
	})
	t.Run("[]bool", func(t *testing.T) {
		assert.Equal(t, []bool{true, false}, Get[[]bool](testdata, "[]bool"))
	})
	t.Run("[]time.Duration", func(t *testing.T) {
		assert.Equal(t, []time.Duration{time.Second}, Get[[]time.Duration](testdata, "[]duration"))
	})
	t.Run("[]any", func(t *testing.T) {
		assert.Equal(t, []any{map[string]any{
			"1": 1,
			"2": 2,
			"3": 3,
		}}, Get[[]any](testdata, "[]any"))
	})
	t.Run("map[string]any", func(t *testing.T) {
		assert.Equal(t, map[string]any{"1": "value1", "2": 1}, Get[map[string]any](testdata, "map[string]any"))
	})
	t.Run("map[string]string", func(t *testing.T) {
		assert.Equal(t, map[string]string{"1": "value"}, Get[map[string]string](testdata, "map[string]string"))
	})
	t.Run("map[string]bool", func(t *testing.T) {
		assert.Equal(t, map[string]bool{"1": true, "2": false}, Get[map[string]bool](testdata, "map[string]bool"))
	})
	t.Run("map[string]int", func(t *testing.T) {
		assert.Equal(t, map[string]int{"1": 1, "2": 2}, Get[map[string]int](testdata, "map[string]int"))
	})
	t.Run("map[string]int64", func(t *testing.T) {
		assert.Equal(t, map[string]int64{"1": 1, "2": 2}, Get[map[string]int64](testdata, "map[string]int64"))
	})
	t.Run("map[string][]string", func(t *testing.T) {
		assert.Equal(t, map[string][]string{"1": {"a", "b", "c"}}, Get[map[string][]string](testdata, "map[string][]string"))
	})
}

func TestMustGet(t *testing.T) {
	var testdata = map[string]any{
		"key": "value",
	}
	assert.Panics(t, func() {
		MustGet[string](testdata, "key1")
	})
	assert.Equal(t, "value", MustGet[string](testdata, "key"))
}

func TestKeys(t *testing.T) {
	var data = map[string]any{
		"name": "test",
		"age":  "18",
		"size": 10,
		"tags": []string{
			"1", "2", "3",
		},
	}
	assert.ElementsMatch(t, []string{"name", "age", "size", "tags"}, Keys(data))
}

func TestValues(t *testing.T) {
	var data = map[string]any{
		"name": "test",
		"age":  "18",
		"size": 10,
		"tags": []string{
			"1", "2", "3",
		},
	}
	assert.ElementsMatch(t, []any{"test", "18", 10, []string{
		"1", "2", "3",
	}}, Values(data))
}

func TestFind(t *testing.T) {
	var data = map[string]any{
		"attributes": map[string]any{
			"name": "test",
			"address": map[string]any{
				"city": "shanghai",
			},
		},
	}
	assert.Equal(t, "test", Find[string](data, "attributes", "name"))
	assert.Equal(t, "shanghai", Find[string](data, "attributes", "address", "city"))
}
