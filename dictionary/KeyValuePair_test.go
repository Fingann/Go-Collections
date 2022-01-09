package dictionary

// import testify
import (
	"testing"
)

func TestKeyValuePair_Key(t *testing.T) {
	kvp := NewKeyValuePair("key", "value")
	if kvp.Key() != "key" {
		t.Errorf("Expected Key() to return 'key', but got %v instead", kvp.Key())
	}
}

func TestKeyValuePair_Value(t *testing.T) {
	kvp := NewKeyValuePair("key", "value")
	if kvp.Value() != "value" {
		t.Errorf("Expected Value() to return 'value', but got %v instead", kvp.Value())
	}
}

func TestNewKeyValuePair(t *testing.T) {
	kvp := NewKeyValuePair("key", "value")
	if kvp.Key() != "key" {
		t.Errorf("Expected Key() to return 'key', but got %v instead", kvp.Key())
	}
	if kvp.Value() != "value" {
		t.Errorf("Expected Value() to return 'value', but got %v instead", kvp.Value())
	}
}
