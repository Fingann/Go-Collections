package internal

import (
	"reflect"
)


type Action[T any] func(object T) error
type Predicate[T any] func(object T) bool
type ComparableFunc[T any, K any] func(object T, object2 K) int
type EquatableFunc[T any] func(object T, object2 T) bool

func NotNil[T any](object T) bool {
	//check if object is nil
	return reflect.ValueOf(object).IsNil()
}

func Equals[T any](object T, object2 T) bool {
	return reflect.DeepEqual(object, object2)
}

type IComparable[T any] interface {
	CompareTo(other T) int
}
type IEquatable[T any] interface {
	Equals(item1 T, item2 T) bool
}
