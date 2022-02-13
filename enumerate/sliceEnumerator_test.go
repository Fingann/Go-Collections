package enumerate

import (
	"testing"
)

func TestSliceEnumerator(t *testing.T) {
	list := []string{"a", "b", "c"}
	enumerator := SliceEnumerator(list)
	if enumerator.Current() != "a" {
		t.Errorf("Expected Current() to return 'a', but got %v instead", enumerator.Current())
	}
	enumerator.Reset()
	if enumerator.MoveNext() != true {
		t.Errorf("Expected MoveNext() to return true, but got %v instead", enumerator.MoveNext())
	}
	if enumerator.Current() != "b" {
		t.Errorf("Expected Current() to return 'b', but got %v instead", enumerator.Current())
	}
	if enumerator.MoveNext() != true {
		t.Errorf("Expected MoveNext() to return true, but got %v instead", enumerator.MoveNext())
	}
	if enumerator.Current() != "c" {
		t.Errorf("Expected Current() to return 'c', but got %v instead", enumerator.Current())
	}
	if enumerator.MoveNext() != false {
		t.Errorf("Expected MoveNext() to return false, but got %v instead", enumerator.MoveNext())
	}
}
