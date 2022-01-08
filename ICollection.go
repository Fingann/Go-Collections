package Collections

import "github.com/Fingann/Go-Collections/enumerate"

type CountableCollection[T any] interface {
	Count() int
}

type SynchronizedCollection[T any] interface {
	GetSyncRoot()
}

type ReadableCollection[T any] interface {
	Get(index int) (T, error)
	Contains(value T) bool
}

type WritableCollection[T any] interface {
	Add(value T) error
	Clear() error
	Remove(value T) error
}

type ReadWriteCollection[T any] interface {
	ReadableCollection[T]
	WritableCollection[T]
}

type Collection[T any] interface {
	CountableCollection[T]
	enumerate.Enumerable[T]
	SynchronizedCollection[T]
	ReadWriteCollection[T]
}
