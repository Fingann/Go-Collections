package dictionary

import (
	Collections "github.com/Fingann/Go-Collections/collections"
)

type IDictionary[TKey comparable, TValue any] interface {
	Collections.ICollection[KeyValuePair[TKey, TValue]]
	AddKeyValue(key TKey, value TValue) error
	ContainsKey(key TKey) bool
	Keys() Collections.ICollection[TKey]
	Values() Collections.ICollection[TValue]
}
