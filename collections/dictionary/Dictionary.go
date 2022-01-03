package dictionary

import (
	"errors"
	"sync"

	Collections "github.com/Fingann/Go-Collections/collections"
	Array "github.com/Fingann/Go-Collections/collections/array"
)

var KeyExistsException = errors.New("An element with the same key already exists in the Dictionary")
var KeyNotFoundException = errors.New("Key does not exists in the Dictionary")

type Dictionary[TKey comparable, TValue any] struct {
	Collections.ICollection[KeyValuePair[TKey, TValue]]
	IDictionary[TKey, TValue]
	dict     map[TKey]TValue
	syncRoot *sync.Mutex
}

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
func (d *Dictionary[TKey, TValue]) GetEnumerator() Collections.IEnumerator[*KeyValuePair[TKey, TValue]] {
	list := make(Array.Array[*KeyValuePair[TKey, TValue]], 0)
	for key, value := range d.dict {
		list.Add(NewKeyValuePair(key, value))
	}
	return Collections.Enumerator(list)

}

// SyncRoot is inherited from ICollection
func (d *Dictionary[TKey, TValue]) SyncRoot() *sync.Mutex {
	return d.syncRoot

}
func (d *Dictionary[TKey, TValue]) Add(pair KeyValuePair[TKey, TValue]) error {
	_, ok := d.dict[pair.Key()]
	if ok {
		return KeyExistsException
	}
	d.dict[pair.Key()] = pair.Value()
	return nil
}

func (d *Dictionary[TKey, TValue]) AddKeyValue(key TKey, value TValue) error {
	_, ok := d.dict[key]
	if ok {
		return KeyExistsException
	}
	d.dict[key] = value
	return nil
}

func (d *Dictionary[TKey, TValue]) ContainsKey(key TKey) bool {
	_, ok := d.dict[key]
	return ok

}

func (d *Dictionary[TKey, TValue]) Remove(key TKey) (bool, error) {
	_, ok := d.dict[key]
	if !ok {
		return false, KeyNotFoundException
	}
	delete(d.dict, key)
	return true, nil

}

func (d *Dictionary[TKey, TValue]) Get(key TKey) (TValue, error) {
	value, ok := d.dict[key]
	if !ok {
		return *new(TValue), KeyNotFoundException
	}
	return value, nil

}

func (d *Dictionary[TKey, TValue]) Keys() Collections.ICollection[TKey] {
	list := Array.New[TKey]()
	for key, _ := range d.dict {
		list.Add(key)
	}
	return list

}

func (d *Dictionary[TKey, TValue]) Values() Collections.ICollection[TValue] {
	list := Array.New[TValue]()
	for _, value := range d.dict {
		list.Add(value)
	}
	return list
}
