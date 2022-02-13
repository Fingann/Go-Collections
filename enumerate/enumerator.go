package enumerate

type Enumerator[T any] struct {
	IEnumerator[T]
}
type IEnumerator[T any] interface {
	Current() T
	Reset()
	MoveNext() bool
}

func NewEnumertor[T any](enumerator IEnumerator[T]) Enumerator[T] {
	return Enumerator[T]{IEnumerator: enumerator}
}
