package Collections

type IEnumerable[T any] interface {
	GetEnumerator() IEnumerator[T]
}

type IEnumerator[T any] interface {
	Current() T
	Resert()
	MoveNext() bool
}

func Enumerator[T any](collection []T) IEnumerator[T] {
	return &enumerator[T]{
		list:    collection,
		current: 0,
	}
}

type enumerator[T any] struct {
	IEnumerator[T]
	list    []T
	current int
}

func (le *enumerator[T]) Current() T {
	return le.list[le.current]
}
func (le *enumerator[T]) Reset() {
	le.current = -1
}
func (le *enumerator[T]) MoveNext() bool {
	le.current = le.current + 1
	return le.current < len(le.list)
}
