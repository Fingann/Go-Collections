package enumerate

type Enumerable[T any] interface {
	GetEnumerator() Enumerator[T]
}
