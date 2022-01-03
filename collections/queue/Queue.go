package queue

import (
	"errors"
	"sync"

	Collections "github.com/Fingann/Go-Collections/collections"
)

var QueueEmptyExcepition = errors.New("Queue is empty")

type Queue[T comparable] struct {
	Collections.ICollection[T]
	items    []T
	syncRoot *sync.Mutex
}

func New[T comparable]() *Queue[T] {
	return &Queue[T]{
		items:    make([]T, 0),
		syncRoot: &sync.Mutex{},
	}
}

// GetEnumerator returns an enumerator that iterates through the List[T]
func (q *Queue[T]) GetEnumerator() Collections.IEnumerator[T] {
	return Collections.Enumerator(q.items)

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

func (q *Queue[T]) Dequeue() (T, error) {
	if len(q.items) <= 0 {
		return *new(T), QueueEmptyExcepition
	}
	front := q.items[len(q.items)-1]
	q.items = q.items[:len(q.items)-1]
	return front, nil
}

func (q *Queue[T]) Enqueue(item T) {
	q.items = append(q.items, item)
}

func (q *Queue[T]) Peek() (T, error) {
	if len(q.items) <= 0 {
		return *new(T), QueueEmptyExcepition
	}
	return q.items[len(q.items)-1], nil
}

func (l *Queue[T]) Count() int {
	return len(l.items)
}
