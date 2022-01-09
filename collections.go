package collection

import (
	"sync"

	"github.com/Fingann/Go-Collections/enumerate"
)

type Countable[T any] interface {
	Count() int
}

type Synchronizable[T any] interface {
	GetSyncRoot() *sync.Mutex
}

type Readable[T any] interface {
	Get(index int) (T, error)
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
	Countable[T]
	enumerate.Enumerable[T]
	Synchronizable[T]
	ReadWriteable[T]
}
