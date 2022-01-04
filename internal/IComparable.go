package internal

type IComparable[T any] interface {
	CompareTo(other T) int
}
