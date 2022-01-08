package dictionary

import (
	"github.com/Fingann/Go-Collections"
	"github.com/Fingann/Go-Collections/list"
)

// IDictionary[TKey, TValue] is the interface that represents a generic dictionary.
type IDictionary[TKey comparable, TValue any] interface {
	collection.Collection[KeyValuePair[TKey, TValue]]
	AddKeyValue(key TKey, value TValue) error
	ContainsKey(key TKey) bool
	Keys() *list.List[TKey]
	Values() *list.List[TValue]
}
