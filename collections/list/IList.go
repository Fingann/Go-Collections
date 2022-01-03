package list

import (
	Collections "github.com/Fingann/Go-Collections/collections"
)

type IList[T any] interface {
	Collections.IEnumerator[T]
	Collections.ICollection[T]
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
