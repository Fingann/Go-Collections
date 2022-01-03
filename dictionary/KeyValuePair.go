package dictionary

// KeyValuePair is a struct that holds a key and a value
type KeyValuePair[TKey any, TValue any] struct {
	key   TKey
	value TValue
}

// NewKeyValuePair returns a new KeyValuePair[TKey, TValue]
func NewKeyValuePair[TKey any, TValue any](key TKey, value TValue) *KeyValuePair[TKey, TValue] {
	return &KeyValuePair[TKey, TValue]{
		key: key, value: value,
	}
}

// Key returns the key of the KeyValuePair[TKey, TValue]
func (kvp *KeyValuePair[TKey, TValue]) Key() TKey {
	return kvp.key
}

// Value returns the value of the KeyValuePair[TKey, TValue]
func (kvp *KeyValuePair[TKey, TValue]) Value() TValue {
	return kvp.value
}
