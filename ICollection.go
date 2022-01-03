package Collections

type ICollection[T any] interface {
	IEnumerable[T]
	// Gets an object that can be used to synchronize access to the Collection.
	Count() int
	Get(index int) (T, error)
	Add(value T) (int, error)
	Clear() error
	Contains(value T) bool
	Remove(value T) error
}
