package enumerate

import (
	"github.com/Fingann/Go-Collections/internal"
)


type Query[T any] struct {
	enumerator IEnumerator[T]
}

func NewQuery[T any](enumerator IEnumerator[T]) *Query[T]{
	return &Query[T]{
		enumerator: enumerator,
	}

}


func (e *Query[T]) Where(p internal.Predicate[T]) *Query[T] {
	filtered := make([]T, 0)
	for e.enumerator.MoveNext() {
		if p(e.enumerator.Current()) {
			filtered = append(filtered, e.enumerator.Current())
		}
	}
	return NewQuery[T](NewSliceEnumerator[T](filtered))
}

func (e *Query[T]) ForEach(action internal.Action[T]) error {
	for e.enumerator.MoveNext() {
		err := action(e.enumerator.Current())
		if err != nil {
			return err
		}

	}
	return nil
}