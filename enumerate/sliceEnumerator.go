package enumerate

func NewSliceEnumerator[T any](items []T) IEnumerator[T] {
	return &SliceEnumerator[T]{
		list:    items,
		current: 0,
	}
}

type SliceEnumerator[T any] struct {
	Enumerator[T]
	list    []T
	current int
}

func (le *SliceEnumerator[T]) Current() T {
	return le.list[le.current]
}
func (le *SliceEnumerator[T]) Reset() {
	le.current = 0
}
func (le *SliceEnumerator[T]) MoveNext() bool {
	le.current = le.current + 1
	return le.current < len(le.list)
}
