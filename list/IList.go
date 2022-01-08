package list

import (
	"github.com/Fingann/Go-Collections"
)

// IList[T] is the interface that represents a generic list.
type IList[T any] interface {
	Collections.Collection[T]
	IsFixedSize() bool
	IsReadOnly() bool
	IndexOf(value T) (int, error)
	Insert(index int, value T) error
	RemoveAt(index int) error
}
