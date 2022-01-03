package Collections

type Comparable interface {
	int64 | int16
}

type IEqualityComparer[T any] interface {
	Equals(item1 T, item2 T) bool
}
