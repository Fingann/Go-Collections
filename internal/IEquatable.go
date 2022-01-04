package internal

type IEquatable[T any] interface {
	Equals(item1 T, item2 T) bool
}
