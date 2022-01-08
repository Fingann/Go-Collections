package enumerate

func SliceEnumerator[T any](collection []T) IEnumerator[T] {
	return &sliceEnumerator[T]{
		list:    collection,
		current: 0,
	}
}

type sliceEnumerator[T any] struct {
	Enumerator[T]
	list    []T
	current int
}

func (le *sliceEnumerator[T]) Current() T {
	return le.list[le.current]
}
func (le *sliceEnumerator[T]) Reset() {
	le.current = 0
}
func (le *sliceEnumerator[T]) MoveNext() bool {
	le.current = le.current + 1
	return le.current < len(le.list)
}
