package enumerate

import "github.com/Fingann/Go-Collections/internal"

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

func (e Enumerator[T]) Where(p internal.Predicate[T]) Enumerator[T] {
	filtered := make([]T, 0)
	for e.MoveNext() {
		if p(e.Current()) {
			filtered = append(filtered, e.Current())
		}
	}
	return NewEnumertor(NewSliceEnumerator(filtered))
}

func (e Enumerator[T]) ForEach(action internal.Action[T]) error {
	for e.MoveNext() {
		err := action(e.Current())
		if err != nil {
			return err
		}

	}
	return nil
}
