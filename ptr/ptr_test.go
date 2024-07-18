package ptr

import (
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func TestPtr(t *testing.T) {
	t.Run("int", func(t *testing.T) {
		assert.Equal(t, 1, *Ptr(1))
		assert.True(t, reflect.ValueOf(Ptr(1)).Type().Kind() == reflect.Ptr)
	})

	t.Run("string", func(t *testing.T) {
		assert.Equal(t, "hello world", *Ptr("hello world"))
		assert.True(t, reflect.ValueOf(Ptr("hello world")).Type().Kind() == reflect.Ptr)
	})

	t.Run("bool", func(t *testing.T) {
		assert.Equal(t, true, *Ptr(true))
		assert.True(t, reflect.ValueOf(Ptr(true)).Type().Kind() == reflect.Ptr)
	})

	t.Run("slice", func(t *testing.T) {
		assert.Equal(t, []int{1, 2, 3}, *Ptr([]int{1, 2, 3}))
		assert.True(t, reflect.ValueOf(Ptr([]int{1, 2, 3})).Type().Kind() == reflect.Ptr)
	})

	t.Run("map", func(t *testing.T) {
		assert.Equal(t, map[string]string{"key": "value"}, *Ptr(map[string]string{"key": "value"}))
		assert.True(t, reflect.ValueOf(Ptr(map[string]string{"key": "value"})).Type().Kind() == reflect.Ptr)
	})

	t.Run("struct", func(t *testing.T) {
		type s struct {
			Value int
		}
		assert.Equal(t, s{Value: 1}, *Ptr(s{
			Value: 1,
		}))
		assert.True(t, reflect.ValueOf(Ptr(s{Value: 1})).Type().Kind() == reflect.Ptr)
	})
}
