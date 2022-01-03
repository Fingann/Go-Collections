package Collection

type IEnumerable[T any] interface {
	GetEnumerator() IEnumerator[T]
}

type IEnumerator[T any] interface {
	Current() T
	Resert()
	MoveNext() bool
}
