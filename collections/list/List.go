package list

import (
	"errors"
	Collections "github.com/Fingann/Go-Collections/collections"
	"sync"

	"github.com/Fingann/Go-Collections/internal"
)

var IndexOutOfRangeException = errors.New("Index was out of range")

type List[T comparable] struct {
	Collections.ICollection[T]
	IList[T]
	items    []T
	syncRoot *sync.Mutex
}

func From[T comparable](list []T) *List[T] {
	return &List[T]{
		items:    list,
		syncRoot: &sync.Mutex{},
	}
}

func New[T comparable]() *List[T] {
	return &List[T]{
		items:    make([]T, 0),
		syncRoot: &sync.Mutex{},
	}
}

// GetEnumerator returns an enumerator that iterates through the List[T]
func (l *List[T]) GetEnumerator() Collections.IEnumerator[T] {
	return Collections.Enumerator(l.items)

}

// SyncRoot is inherited from ICollection
func (l *List[T]) SyncRoot() *sync.Mutex {
	return l.syncRoot

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

	return From(l.items[index : index+count]), nil

}

func (l *List[T]) Add(value T) (int, error) {
	l.items = append(l.items, value)
	return len(l.items), nil

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
		if s == value {
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
	enumerator := l.GetEnumerator()
	for {
		if predicate(enumerator.Current()) {
			return enumerator.Current(), nil
		}
		if !enumerator.MoveNext() {
			return *new(T), errors.New("Could not find item in collection")
		}
	}
}
func (l *List[T]) FindAll(predicate internal.Predicate[T]) []T {
	enumerator := l.GetEnumerator()
	items := make([]T, 0)
	for {
		if predicate(enumerator.Current()) {
			items = append(items, enumerator.Current())
		}
		if !enumerator.MoveNext() {
			return items
		}
	}
}

func (l *List[T]) ForEach(action internal.Action[T]) error {
	enumerator := l.GetEnumerator()
	for {
		action(enumerator.Current())

		if !enumerator.MoveNext() {
			return errors.New("Could not find item in collection")
		}
	}
}

func (l *List[T]) Count() int {
	return len(l.items)
}
