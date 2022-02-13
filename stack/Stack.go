package stack

import (
	"errors"
	"sync"

	Collections "github.com/Fingann/Go-Collections"
)

var StackEmptyExcepition = errors.New("Stack is empty")

// Stack is a first-in, first-out collection of objects.
type Stack[T any] struct {
	Collections.Collection[T]
	items    []T
	syncRoot *sync.Mutex
}

// New returns a new Stack[T]
func New[T comparable]() *Stack[T] {
	return &Stack[T]{
		items:    make([]T, 0),
		syncRoot: &sync.Mutex{},
	}
}

// SyncRoot is inherited from ICollection
func (q *Stack[T]) SyncRoot() *sync.Mutex {
	return q.syncRoot

}

// Pop removes and returns the object at the top of the Stack.
func (q *Stack[T]) Pop() (T, error) {
	if len(q.items) <= 0 {
		return *new(T), StackEmptyExcepition
	}
	front := q.items[len(q.items)-1]
	q.items = q.items[:len(q.items)-1]
	return front, nil
}

// Push adds an object to the top of the Stack.
func (q *Stack[T]) Push(item T) {
	q.items = append(q.items, item)
}

// Peek returns the object at the top of the Stack without removing it.
func (q *Stack[T]) Peek() (T, error) {
	if len(q.items) <= 0 {
		return *new(T), StackEmptyExcepition
	}
	return q.items[len(q.items)-1], nil
}

// Count returns the number of elements in the Stack.
func (l *Stack[T]) Count() int {
	return len(l.items)
}
