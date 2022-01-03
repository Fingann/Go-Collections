package Collections

import (
	"errors"
	Collections "github.com/Fingann/Go-Collections/collections"
	"reflect"
)

var IndexOutOfRangeException = errors.New("Index was out of range")

// Array is a list of items.
type Array[T any] []T

// From creates a new Array[T] from a slice of items.
func From[T any](list []T) Array[T] {
	return Array[T](list)
}

// New creates a new empty Array[T]
func New[T any]() Array[T] {
	return make(Array[T], 0)
}

// GetEnumerator returns an enumerator that iterates through the Array[T]
func (l Array[T]) GetEnumerator() Collections.IEnumerator[T] {
	return Collections.Enumerator(l)
}

// Count returns the number of items in the Array.
func (l Array[T]) Count() int {
	return len(l)
}

// Get returns the item at the specified index.
func (l Array[T]) Get(index int) (T, error) {
	if index < 0 && index >= len(l) {
		return *new(T), IndexOutOfRangeException
	}

	return l[index], nil
}

// Add adds an item to the Array.
func (l Array[T]) Add(value T) (int, error) {
	l = append(l, value)
	return len(l), nil

}

// Clear removes all items from the Array. Without reducing the Capacity of the array.
func (l Array[T]) Clear() error {
	l = l[:0]
	return nil

}

// Contains determines whether an element is in the List.
func (l Array[T]) Contains(value T) bool {
	for _, item := range l {
		if reflect.DeepEqual(item, value) {
			return true
		}
	}
	return false
}

// Remove removes the first occurrence of an item from the Array.
func (l Array[T]) Remove(value T) error {
	for index, item := range l {
		if reflect.DeepEqual(item, value) {
			l = append(l[:index], l[index+1:]...)
			return nil
		}
	}
	return errors.New("Could not find item in collection")

}
