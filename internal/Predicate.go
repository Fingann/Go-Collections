package internal

type Predicate[T any] func(object T) bool
