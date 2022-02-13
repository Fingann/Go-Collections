package enumerate

type Enumerable[T any] interface {
	GetEnumerable() Enumerator[T]
}
