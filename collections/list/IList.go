package list

import (
	Collections "github.com/Fingann/Go-Collections/collections"
)

type IList[T any] interface {
	Collections.ICollection[T]
	IsFixedSize() bool
	IsReadOnly() bool
	IndexOf(value T) (int, error)
	Insert(index int, value T) error
	RemoveAt(index int) error
}
