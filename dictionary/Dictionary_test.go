package dictionary

import (
	"testing"
)

func VerifyEntry(t *testing.T, d *Dictionary[string, int], key string, expected int) {
	value, err := d.Get(key)
	if err != nil {
		t.Errorf("Expected Get() to return a value, but got an error instead")
	}
	if value != 1 {
		t.Errorf("Expected Get() to return %v, but got %v instead", expected, value)
	}
}

func TestNew(t *testing.T) {
	d := New[string, int]()
	if d == nil {
		t.Errorf("Expected New() to return a dictionary, but got nil instead")
	}
}

func TestFrom(t *testing.T) {
	d := From(map[string]int{
		"key": 1,
	})
	if d == nil {
		t.Errorf("Expected From() to return a dictionary, but got nil instead")
	}
	VerifyEntry(t, d, "key", 1)
}

func TestGetEnumerable(t *testing.T) {
	d := From(map[string]int{
		"key": 1,
	})

	enumerable := d.GetEnumerable()
	if enumerable == nil {
		t.Errorf("Expected GetEnumerable() to return an enumerable, but got nil instead")
	}
}

func TestGetSyncRoot(t *testing.T) {
	d := From(map[string]int{
		"key": 1,
	})

	syncRoot := d.GetSyncRoot()
	if syncRoot == nil {
		t.Errorf("Expected GetSyncRoot() to return a sync root, but got nil instead")
	}
}

func TestAdd(t *testing.T) {
	d := New[string, int]()
	err := d.Add(KeyValuePair[string, int]{"key", 1})
	if err != nil {
		t.Errorf("Expected Add() to return nil, but got %v instead", err)
	}
	VerifyEntry(t, d, "key", 1)
}

func TestAddKeyValue(t *testing.T) {
	d := New[string, int]()
	err := d.AddKeyValue("key", 1)
	if err != nil {
		t.Errorf("Expected AddKeyValue() to return nil, but got %v instead", err)
	}
	VerifyEntry(t, d, "key", 1)
}

func TestContainsKey(t *testing.T) {
	d := From(map[string]int{
		"key": 1,
	})
	if !d.ContainsKey("key") {
		t.Errorf("Expected ContainsKey() to return true, but got false instead")
	}
}

func TestRemove(t *testing.T) {
	d := From(map[string]int{
		"key": 1,
	})
	ok, err := d.Remove("key")
	if err != nil {
		t.Errorf("Expected Remove() to return nil, but got %v instead", err)
	}
	if !ok {
		t.Errorf("Expected Remove() to return true, but got false instead")
	}
	if d.ContainsKey("key") {
		t.Errorf("Expected Remove() to remove the key, but the key still exists")
	}
}

func TestGet(t *testing.T) {
	d := From(map[string]int{
		"key": 1,
	})
	value, err := d.Get("key")
	if err != nil {
		t.Errorf("Expected Get() to return a value, but got an error instead")
	}
	if value != 1 {
		t.Errorf("Expected Get() to return %v, but got %v instead", 1, value)
	}
}

func TestKeys(t *testing.T) {
	d := From(map[string]int{
		"key": 1,
	})
	keys := d.Keys()
	if keys == nil {
		t.Errorf("Expected Keys() to return a list, but got nil instead")
	}
	if keys.Count() != 1 {
		t.Errorf("Expected Keys() to return a list with 1 element, but got %v instead", keys.Count())
	}
	first, _ := keys.Get(0)
	if first != "key" {
		t.Errorf("Expected Keys() to return a list with the key 'key', but got %v instead", first)
	}
}

func TestValues(t *testing.T) {
	d := From(map[string]int{
		"key": 1,
	})
	values := d.Values()
	if values == nil {
		t.Errorf("Expected values() to return a list, but got nil instead")
	}
	if values.Count() != 1 {
		t.Errorf("Expected values() to return a list with 1 element, but got %v instead", values.Count())
	}
	first, _ := values.Get(0)
	if first != 1 {
		t.Errorf("Expected values() to return a list with the Value '1', but got %v instead", first)
	}
}

func TestForEach(t *testing.T) {
	d := From(map[string]int{
		"key": 1,
	})
	d.ForEach(func(kv KeyValuePair[string, int]) {
		if kv.key != "key" {
			t.Errorf("Expected ForEach() to iterate over the key 'key', but got %v instead", kv.key)
		}
		if kv.value != 1 {
			t.Errorf("Expected ForEach() to iterate over the value '1', but got %v instead", kv.value)
		}
	})
}
