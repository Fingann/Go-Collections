package dictionary

import (
	"github.com/Fingann/Go-Collections"
)

// IDictionary[TKey, TValue] is the interface that represents a generic dictionary.
type IDictionary[TKey comparable, TValue any] interface {
	Collections.ICollection[KeyValuePair[TKey, TValue]]
	AddKeyValue(key TKey, value TValue) error
	ContainsKey(key TKey) bool
	Keys() Collections.ICollection[TKey]
	Values() Collections.ICollection[TValue]
}
