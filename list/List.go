package list

import (
	"errors"
	"reflect"
	"sync"

	"github.com/Fingann/Go-Collections/enumerate"
	"github.com/Fingann/Go-Collections/internal"
)

var _ enumerate.Enumerable[string] = New[string]()

var ErrIndexOutOfRange = errors.New("index was out of range")
var ErrItemNotFound = errors.New("item not found")

type List[T any] struct {
	IList[T]
	items    []T
	syncRoot *sync.Mutex
}


func From[T any](list ...T) *List[T] {
	return &List[T]{
		items:    list,
		syncRoot: &sync.Mutex{},
	}
}

func New[T any]() *List[T] {
	return WithLengthCapacity[T](0, 0)
}

func WithLengthCapacity[T any](length int, capacity int) *List[T] {
	items:= make([]T, length, capacity)
	return &List[T]{
		items:    items,
		syncRoot: &sync.Mutex{},
	}
}

func (l *List[T]) GetSyncRoot() *sync.Mutex {
	return l.syncRoot
}

// GetEnumerable returns an enumerator that iterates through the List[T]
func (l *List[T]) GetEnumerator() enumerate.Enumerator[T] {
	return enumerate.NewEnumertor(enumerate.NewSliceEnumerator(l.items))

}

func (l *List[T]) IsFixedSize() bool {
	return false

}
func (l *List[T]) IsReadOnly() bool {
	return false

}
func (l *List[T]) Get(index int) (T, error) {
	if index < 0 || index >= len(l.items) {
		return *new(T), ErrIndexOutOfRange
	}

	return l.items[index], nil
}
func (l *List[T]) GetIndex(index int) (T, error) {
	if index < 0 || index >= len(l.items) {
		return *new(T), ErrIndexOutOfRange
	}

	return l.items[index], nil
}
func (l *List[T]) GetRange(index int, count int) (*List[T], error) {
	if index < 0 || index+count >= len(l.items) {
		return new(List[T]), ErrIndexOutOfRange
	}

	return From(l.items[index : index+count]...), nil
}

func (l *List[T]) Add(value T) *List[T] {
	l.items = append(l.items, value)
	return l
}

func (l *List[T]) Set(index int, value T) error {
	if index < 0 && index >= len(l.items) {
		return ErrIndexOutOfRange
	}
	l.items[index] = value
	return nil
}

func (l *List[T]) AddRange(collection enumerate.Enumerable[T]) *List[T] {
	enumerator := collection.GetEnumerator()
	for {
		l.items = append(l.items, enumerator.Current())
		if !enumerator.MoveNext() {
			break
		}
	}
	return l
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
	return -1, ErrItemNotFound

}

// Insert inserts an item at the specified index.
func (l *List[T]) Insert(index int, value T) error {
	if index < 0 || index >= len(l.items) {
		return ErrIndexOutOfRange
	}
	l.items[index] = value
	return nil

}

// Remove removes the first occurrence of a specific object from the List[T].
func (l *List[T]) Remove(value T) error {
	index, err := l.IndexOf(value)
	if err != nil {
		return err
	}
	return l.RemoveAt(index)

}

// RemoveAt removes the element at the specified index of the List[T].
func (l *List[T]) RemoveAt(index int) error {
	l.items = append(l.items[:index], l.items[index+1:]...)
	return nil

}

//Find returns the first element in the list that satisfies the predicate.
func (l *List[T]) Find(predicate internal.Predicate[T]) (T, error) {
	for _, v := range l.items {
		if predicate(v) {
			return v, nil
		}
	}

	return *new(T), errors.New("could not find item in collection")
}

func (l *List[T]) FindAll(predicate internal.Predicate[T]) *List[T] {
	items := New[T]()
	for _, v := range l.items {
		if predicate(v) {
			items.Add(v)
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

func (l *List[T]) ToSlice() []T {
	return l.items
}
