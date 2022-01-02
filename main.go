package main

import (
	"errors"
	"fmt"
	"sync"
)

func main() {
	lst := NewList([]string{"hello,", " world"})
	val, _:= lst.Get(0)
	fmt.Print(val)
	val, _= lst.Get(1)
	fmt.Println(val)


	list := ArrayList[string]{"hello", "world"}
	enumerator := list.GetEnumerator()
	fmt.Println(enumerator.Current())
	enumerator.MoveNext()
	fmt.Println(enumerator.Current())

}

type ArrayList[E any] []E

func (l ArrayList[T]) GetEnumerator() IEnumerator[T] {
	return &ListEnumerator[T]{
		list: l,
	}
}

type ListEnumerator[T any] struct {
	IEnumerator[T]
	list    ArrayList[T]
	current int
}

func (le *ListEnumerator[T]) Current() T {
	return le.list[le.current]
}
func (le *ListEnumerator[T]) Reset() {
	le.current = -1
}
func (le *ListEnumerator[T]) MoveNext() bool {
	le.current = le.current + 1
	return le.current < len(le.list)
}

/////////// IEnumerable / Enumerator ////////

type IEnumerable[T any] interface {
	GetEnumerator() IEnumerator[T]
}

type IEnumerator[T any] interface {
	Current() T
	Resert()
	MoveNext() bool
}

/////////// ICollection ////////

var ArgumentOutOfRangeException = errors.New("Argument was out of range")

type ICollection[T any] interface {
	IEnumerable[T]
	// Gets an object that can be used to synchronize access to the Collection.
	SyncRoot() *sync.Mutex
	//CopyTo(array)
}

/////////// IList ////////

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

func NewList[T comparable](list []T) List[T]{
	return List[T]{
	items: list,
	syncRoot: &sync.Mutex{},
	}
}

// GetEnumerator is inherited from IEnumerable
func (l List[T]) GetEnumerator() IEnumerator[T] {
	return &ListEnumerator[T]{
		list: l.items,
	}
}

// SyncRoot is inherited from ICollection
func (l List[T]) SyncRoot() *sync.Mutex {
	return l.syncRoot

}

func (l List[T]) IsFixedSize() bool {
	return false

}
func (l List[T]) IsReadOnly() bool {
	return false

}
func (l List[T]) Get(index int) (T, error) {
	if index < 0 && index > len(l.items) {
		var notfound T
		return notfound, ArgumentOutOfRangeException
	}

	return l.items[index], nil

}
func (l List[T]) Add(value T) (int, error) {
	l.items = append(l.items, value)
	return len(l.items), nil

}
func (l List[T]) Clear() error {
	l.items = l.items[:0]
	return nil

}
func (l List[T]) Contains(value T) bool {
	_, err := l.IndexOf(value)
	return err == nil
}

func (l List[T]) IndexOf(value T) (int, error) {
	for i, s := range l.items {
		if s == value {
			return i, nil
		}
	}
	return -1, errors.New("Value is not in index")

}
func (l List[T]) Insert(index int, value T) error {
	if index < 0 && index > len(l.items) {
		return ArgumentOutOfRangeException
	}
	l.items[index] = value
	return nil

}
func (l List[T]) Remove(value T) error {
	index, err := l.IndexOf(value)
	if err != nil {
		return err
	}
	return l.RemoveAt(index)

}

func (l List[T]) RemoveAt(index int) error {
	l.items = append(l.items[:index], l.items[index+1:]...)
	return nil

}

