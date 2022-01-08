package internal

type Action[T any] func(object T)
type Predicate[T any] func(object T) bool
type ComparableFunc[T any, K any] func(object T, object2 K) int
type EquatableFunc[T any] func(object T, object2 T) bool

type IComparable[T any] interface {
	CompareTo(other T) int
}
type IEquatable[T any] interface {
	Equals(item1 T, item2 T) bool
}
