package list

import (
	"errors"
	"reflect"
	"sync"

	Collections "github.com/Fingann/Go-Collections"
	"github.com/Fingann/Go-Collections/internal"
)

var IndexOutOfRangeException = errors.New("Index was out of range")

type List[T any] struct {
	Collections.ICollection[T]
	IList[T]
	items    []T
	syncRoot *sync.Mutex
}

func From[T any](list []T) *List[T] {
	return &List[T]{
		items:    list,
		syncRoot: &sync.Mutex{},
	}
}

func New[T any]() *List[T] {
	return WithLengthCapacity[T](0, 0)
}

func WithLengthCapacity[T any](length int, capacity int) *List[T] {
	return &List[T]{
		items:    make([]T, length, capacity),
		syncRoot: &sync.Mutex{},
	}
}

// GetEnumerator returns an enumerator that iterates through the List[T]
func (l *List[T]) GetEnumerator() Collections.IEnumerator[T] {
	return Collections.Enumerator(l.items)

}

func (l *List[T]) IsFixedSize() bool {
	return false

}
func (l *List[T]) IsReadOnly() bool {
	return false

}
func (l *List[T]) Get(index int) (T, error) {
	if index < 0 && index >= len(l.items) {
		return *new(T), IndexOutOfRangeException
	}

	return l.items[index], nil
}

func (l *List[T]) GetRange(index int, count int) (*List[T], error) {
	if index < 0 && index+count >= len(l.items) {
		return new(List[T]), IndexOutOfRangeException
	}

	return From[T](l.items[index : index+count]), nil
}

func (l *List[T]) Add(value T) error {
	l.items = append(l.items, value)
	return nil
}

func (l *List[T]) Set(index int, value T) error {
	if index < 0 && index >= len(l.items) {
		return IndexOutOfRangeException
	}
	l.items[index] = value
	return nil
}

func (l *List[T]) AddRange(collection Collections.IEnumerable[T]) error {
	enumerator := collection.GetEnumerator()
	for {
		l.items = append(l.items, enumerator.Current())
		if !enumerator.MoveNext() {
			break
		}
	}
	return nil
}

func (l *List[T]) Clear() error {
	l.items = l.items[:0]
	return nil

}

// Contains determines whether an element is in the List.
func (l *List[T]) Contains(value T) bool {
	_, err := l.IndexOf(value)
	return err == nil
}

func (l *List[T]) IndexOf(value T) (int, error) {
	for i, s := range l.items {
		if reflect.DeepEqual(s, value) {
			return i, nil
		}
	}
	return -1, errors.New("Value is not in index")

}
func (l *List[T]) Insert(index int, value T) error {
	if index < 0 && index > len(l.items) {
		return IndexOutOfRangeException
	}
	l.items[index] = value
	return nil

}
func (l *List[T]) Remove(value T) error {
	index, err := l.IndexOf(value)
	if err != nil {
		return err
	}
	return l.RemoveAt(index)

}

func (l *List[T]) RemoveAt(index int) error {
	l.items = append(l.items[:index], l.items[index+1:]...)
	return nil

}

func (l *List[T]) Find(predicate internal.Predicate[T]) (T, error) {
	for _, v := range l.items {
		if predicate(v) {
			return v, nil
		}
	}

	return *new(T), errors.New("Could not find item in collection")
}

func (l *List[T]) FindAll(predicate internal.Predicate[T]) []T {
	items := make([]T, 0)
	for _, v := range l.items {
		if predicate(v) {
			items = append(items, v)
		}
	}
	return items
}

func (l *List[T]) ForEach(action internal.Action[T]) error {
	for _, v := range l.items {
		action(v)
	}
	return nil
}

func (l *List[T]) Count() int {
	return len(l.items)
}
