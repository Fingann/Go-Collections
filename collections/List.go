package Collections

import (
	"errors"
	"sync"

	"github.com/Fingann/Go-Collections/internal"
)

type IList[T any] interface {
	IEnumerator[T]
	ICollection[T]
	IsFixedSize() bool
	IsReadOnly() bool
	Get(index int) (T, error)
	Add(value T) (int, error)
	Clear() error
	Contains(value T) bool
	IndexOf(value T) (int, error)
	Insert(index int, value T) error
	Remove(value T) error
	RemoveAt(index int) error
}

/////////// List Base ////////

type List[T comparable] struct {
	ICollection[T]
	IList[T]
	items    []T
	syncRoot *sync.Mutex
}

func NewList[T comparable](list []T) *List[T] {
	return &List[T]{
		items:    list,
		syncRoot: &sync.Mutex{},
	}
}

// GetEnumerator returns an enumerator that iterates through the List[T]
func (l *List[T]) GetEnumerator() IEnumerator[T] {
	return Enumerator(l.items)

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

	return NewList(l.items[index : index+count]), nil

}

func (l *List[T]) Add(value T) (int, error) {
	l.items = append(l.items, value)
	return len(l.items), nil

}

func (l *List[T]) AddRange(collection IEnumerable[T]) error {
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
		l.items = append(l.items, enumerator.Current())
		if !enumerator.MoveNext() {
			return *new(T), errors.New("Could not find item in collection")
		}
	}
}
