package dictionary

import (
	"errors"
	"sync"

	"github.com/Fingann/Go-Collections/enumerate"
	"github.com/Fingann/Go-Collections/internal"
	"github.com/Fingann/Go-Collections/list"
)

var KeyExistsException = errors.New("An element with the same key already exists in the Dictionary")
var KeyNotFoundException = errors.New("Key does not exists in the Dictionary")

// Dictionary represents a map of key-value pairs.
type Dictionary[TKey comparable, TValue any] struct {
	IDictionary[TKey, TValue]
	enumerate.Enumerable[KeyValuePair[TKey, TValue]]
	dict     map[TKey]TValue
	syncRoot *sync.Mutex
}

// From creates a new Dictionary[TKey, TValue] from a map[TKey]TValue.
func From[TKey comparable, TValue any](dict map[TKey]TValue) *Dictionary[TKey, TValue] {
	return &Dictionary[TKey, TValue]{
		dict:     dict,
		syncRoot: &sync.Mutex{},
	}
}

func New[TKey comparable, TValue any]() *Dictionary[TKey, TValue] {
	return &Dictionary[TKey, TValue]{
		dict:     make(map[TKey]TValue),
		syncRoot: &sync.Mutex{},
	}
}

// GetEnumerator returns an enumerator that iterates through the List[T]
func (d *Dictionary[TKey, TValue]) GetEnumerator() *enumerate.Enumerator[KeyValuePair[TKey, TValue]] {
	list := make([]KeyValuePair[TKey, TValue], 0, len(d.dict))
	for key, value := range d.dict {
		list = append(list, KeyValuePair[TKey, TValue]{key, value})
	}
	return enumerate.NewEnumertor(enumerate.SliceEnumerator(list))

}

// SyncRoot is inherited from ICollection
func (d *Dictionary[TKey, TValue]) SyncRoot() *sync.Mutex {
	return d.syncRoot

}

// Add adds a key-value pair to the Dictionary[TKey, TValue].
func (d *Dictionary[TKey, TValue]) Add(pair KeyValuePair[TKey, TValue]) error {
	_, ok := d.dict[pair.Key()]
	if ok {
		return KeyExistsException
	}
	d.dict[pair.Key()] = pair.Value()
	return nil
}

// AddKeyValue adds a key-value pair to the Dictionary[TKey, TValue].
func (d *Dictionary[TKey, TValue]) AddKeyValue(key TKey, value TValue) error {
	_, ok := d.dict[key]
	if ok {
		return KeyExistsException
	}
	d.dict[key] = value
	return nil
}

// ContainsKey checks if the Dictionary[TKey, TValue] contains a key.
func (d *Dictionary[TKey, TValue]) ContainsKey(key TKey) bool {
	_, ok := d.dict[key]
	return ok

}

// Remove removes a key-value pair from the Dictionary[TKey, TValue].
// It returns true if the key-value pair is successfully removed; otherwise, false.
func (d *Dictionary[TKey, TValue]) Remove(key TKey) (bool, error) {
	_, ok := d.dict[key]
	if !ok {
		return false, KeyNotFoundException
	}
	delete(d.dict, key)
	return true, nil

}

// Get returns the value for a key.
// It returns the value, or an error if the key is not found.
func (d *Dictionary[TKey, TValue]) Get(key TKey) (TValue, error) {
	value, ok := d.dict[key]
	if !ok {
		return *new(TValue), KeyNotFoundException
	}
	return value, nil

}

// Keys returns a collection of keys in the Dictionary[TKey, TValue].
func (d *Dictionary[TKey, TValue]) Keys() *list.List[TKey] {
	keys := make([]TKey, 0, len(d.dict))
	for key, _ := range d.dict {
		keys = append(keys, key)
	}
	return list.From(keys)

}

// Values returns a collection of values in the Dictionary[TKey, TValue].
func (d *Dictionary[TKey, TValue]) Values() *list.List[TValue] {
	values := make([]TValue, 0, len(d.dict))
	for _, value := range d.dict {
		values = append(values, value)
	}
	return list.From(values)
}

// Foreach calls the specified action for each key-value pair in the Dictionary[TKey, TValue].
func (l *Dictionary[TKey, TValue]) ForEach(action internal.Action[KeyValuePair[TKey, TValue]]) error {
	for k, v := range l.dict {
		action(KeyValuePair[TKey, TValue]{k, v})
	}
	return nil
}
