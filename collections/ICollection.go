package Collections

import "sync"

type ICollection[T any] interface {
	IEnumerable[T]
	// Gets an object that can be used to synchronize access to the Collection.
	SyncRoot() *sync.Mutex
	Count() int
}
