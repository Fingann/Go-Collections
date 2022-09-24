package collection

import (
	"github.com/Fingann/Go-Collections/enumerate"
)

type Countable[T any] interface {
	Count() int
}



type Readable[T any] interface {
	Get(item T) (T, error)
	Contains(value T) bool
}

type Writable[T any] interface {
	Add(value T) error
	Clear() error
	Remove(value T) error
}

type ReadWriteable[T any] interface {
	Readable[T]
	Writable[T]
}

type Collection[T any] interface {
	enumerate.Enumerable[T]
	ReadWriteable[T]
	ToSlice() []T
}
