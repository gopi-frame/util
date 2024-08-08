package kv

import (
	"fmt"
	"reflect"
	"time"

	"github.com/gopi-frame/exception"
	"github.com/spf13/cast"
)

// Get gets value from map and convert it to target type T
// if key is not found in the given map, the zero value of type T will be returned
// if value can not be converted to type T, the zero value of type T will be returned
func Get[T any, K comparable](data map[K]any, key K) T {
	value, _ := GetE[T](data, key)
	return value
}

// MustGet must get value map and convert it to target type T
// if key is not found in the given map, an [exception.UndefinedIndexException] will be raised
// if value can't be converted to type T, a [exception.TypeException] will be raised
func MustGet[T any, K comparable](data map[K]any, key K) T {
	if value, err := GetE[T](data, key); err != nil {
		panic(err)
	} else {
		return value
	}
}

// GetE get value from map and convert it to target type T
// if key is not found in the given map, the zero value of type T and an [exception.UndefinedIndexException] will be returned
// if value can not be converted to type T, the zero value of type T and a [exception.TypeException] will be returned
func GetE[T any, K comparable](data map[K]any, key K) (T, error) {
	value, ok := data[key]
	if !ok {
		return *new(T), exception.NewUndefinedIndexException(fmt.Sprintf("%v", key))
	}
	if value, ok := value.(T); ok {
		return value, nil
	}
	var convert = func(v any) T {
		return reflect.ValueOf(v).Interface().(T)
	}
	var v any = *new(T)
	switch v.(type) {
	case string:
		v, err := cast.ToStringE(value)
		return convert(v), err
	case int:
		v, err := cast.ToIntE(value)
		return convert(v), err
	case uint:
		v, err := cast.ToUintE(value)
		return convert(v), err
	case int8:
		v, err := cast.ToInt8E(value)
		return convert(v), err
	case uint8:
		v, err := cast.ToUint8E(value)
		return convert(v), err
	case int16:
		v, err := cast.ToInt16E(value)
		return convert(v), err
	case uint16:
		v, err := cast.ToUint16E(value)
		return convert(v), err
	case int32:
		v, err := cast.ToInt32E(value)
		return convert(v), err
	case uint32:
		v, err := cast.ToUint32E(value)
		return convert(v), err
	case int64:
		v, err := cast.ToInt64E(value)
		return convert(v), err
	case uint64:
		v, err := cast.ToUint64E(value)
		return convert(v), err
	case float32:
		v, err := cast.ToFloat32E(value)
		return convert(v), err
	case float64:
		v, err := cast.ToFloat64E(value)
		return convert(v), err
	case bool:
		v, err := cast.ToBoolE(value)
		return convert(v), err
	case time.Duration:
		v, err := cast.ToDurationE(value)
		return convert(v), err
	case time.Time:
		v, err := cast.ToTimeE(value)
		return convert(v), err
	case []int:
		v, err := cast.ToIntSliceE(value)
		return convert(v), err
	case []string:
		v, err := cast.ToStringSliceE(value)
		return convert(v), err
	case []bool:
		v, err := cast.ToBoolSliceE(value)
		return convert(v), err
	case []time.Duration:
		v, err := cast.ToDurationSliceE(value)
		return convert(v), err
	case []any:
		v, err := cast.ToSliceE(value)
		return convert(v), err
	case map[string]any:
		v, err := cast.ToStringMapE(value)
		return convert(v), err
	case map[string]string:
		v, err := cast.ToStringMapStringE(value)
		return convert(v), err
	case map[string]bool:
		v, err := cast.ToStringMapBoolE(value)
		return convert(v), err
	case map[string]int:
		v, err := cast.ToStringMapIntE(value)
		return convert(v), err
	case map[string]int64:
		v, err := cast.ToStringMapInt64E(value)
		return convert(v), err
	case map[string][]string:
		v, err := cast.ToStringMapStringSliceE(value)
		return convert(v), err
	}
	return *new(T), exception.NewTypeException(fmt.Sprintf("value can't convert to type \"%T\"", *new(T)))
}

func Keys[T comparable](data map[T]any) (keys []T) {
	for key := range data {
		keys = append(keys, key)
	}
	return
}

func Values[T any](data map[string]T) (values []T) {
	for _, value := range data {
		values = append(values, value)
	}
	return
}

func Find[T any](data map[string]any, keys ...string) T {
	value, _ := FindE[T](data, keys...)
	return value
}

func FindE[T any](data map[string]any, keys ...string) (T, error) {
	for index, key := range keys {
		if index < len(keys)-1 {
			value, err := GetE[map[string]any](data, key)
			if err != nil {
				return *new(T), err
			}
			data = value
		} else {
			value, err := GetE[T](data, key)
			if err != nil {
				return *new(T), err
			}
			return value, nil
		}
	}
	return *new(T), nil
}
