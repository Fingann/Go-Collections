package Collections

type IEnumerator[T any] interface {
	Current() T
	Resert()
	MoveNext() bool
}
