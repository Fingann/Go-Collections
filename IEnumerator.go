package Collections

// Ienumerator[T] is an interface that defines the methods of an enumerator.
type IEnumerator[T any] interface {
	Current() T
	Resert()
	MoveNext() bool
}
