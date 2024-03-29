package queue

import (
	"errors"
	"sync"

	"github.com/Fingann/Go-Collections"
)

var QueueEmptyExcepition = errors.New("Queue is empty")

// Queue is a first-in, first-out collection of objects.
type Queue[T comparable] struct {
	collection.Collection[T]
	items    []T
	syncRoot *sync.Mutex
}

// New returns a new Queue[T]
func New[T comparable]() *Queue[T] {
	return &Queue[T]{
		items:    make([]T, 0),
		syncRoot: &sync.Mutex{},
	}
}

// SyncRoot is inherited from ICollection
func (q *Queue[T]) SyncRoot() *sync.Mutex {
	return q.syncRoot

}

// Contains determines whether an element is in the Queue.
func (q *Queue[T]) Contains(value T) bool {
	for _, s := range q.items {
		if s == value {
			return true
		}
	}
	return false
}

// Dequeue removes and returns the object at the beginning of the Queue.
func (q *Queue[T]) Dequeue() (T, error) {
	if len(q.items) <= 0 {
		return *new(T), QueueEmptyExcepition
	}
	front := q.items[0]
	q.items = q.items[1:]
	return front, nil
}

// Enqueue adds an object to the end of the Queue.
func (q *Queue[T]) Enqueue(item T) *Queue[T] {
	q.items = append(q.items, item)
	return q
}

// Peek returns the object at the beginning of the Queue without removing it.
func (q *Queue[T]) Peek() (T, error) {
	if len(q.items) <= 0 {
		return *new(T), QueueEmptyExcepition
	}
	return q.items[0], nil
}

// Count returns the number of elements in the Queue.
func (l *Queue[T]) Count() int {
	return len(l.items)
}

func (q *Queue[T]) ToSlice() []T {
	return q.items
}