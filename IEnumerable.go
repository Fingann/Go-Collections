package Collections

// IEnumerable[T] is an interface that defines a method for retreiving an Enumerator for a given collection
type IEnumerable[T any] interface {
	GetEnumerator() IEnumerator[T]
}

func ToSlice[T IEnumerable[T]](collection T) []T {
	slice := make([]T, 0)
	enumerator := collection.GetEnumerator()
	for {
		slice = append(slice, enumerator.Current())
		if !enumerator.MoveNext() {
			return slice
		}
	}
}