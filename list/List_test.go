package list

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func VerifyEntry(t *testing.T, l *List[int], index int, expected int) {
	value, err := l.Get(index)
	if err != nil {
		t.Errorf("Expected Get() to return a value, but got an error instead")
	}
	if value != expected {
		t.Errorf("Expected Get() to return %v, but got %v instead", expected, value)
	}
}

func TestFrom(t *testing.T) {
	d := From([]int{1, 2, 3})
	if d == nil {
		t.Errorf("Expected From() to return a List, but got nil instead")
	}
	VerifyEntry(t, d, 0, 1)
	VerifyEntry(t, d, 1, 2)
	VerifyEntry(t, d, 2, 3)
}

func TestNew(t *testing.T) {
	d := New[int]()
	if d == nil {
		t.Errorf("Expected New() to return a List, but got nil instead")
	}
}
func TestGetSyncRoot(t *testing.T) {
	d := From([]int{1, 2, 3})

	syncRoot := d.GetSyncRoot()
	if syncRoot == nil {
		t.Errorf("Expected GetSyncRoot() to return a sync root, but got nil instead")
	}
}

func TestWithLengthCapacity(t *testing.T) {
	d := WithLengthCapacity[int](1, 3)
	if d == nil {
		t.Errorf("Expected WithLengthCapacity() to return a List, but got nil instead")
	}
	_, err := d.Get(0)
	if err != nil {
		t.Errorf("Expected Get() to return a value, but got an error instead")
	}
	d2 := WithLengthCapacity[int](0, 3)
	if d == nil {
		t.Errorf("Expected WithLengthCapacity() to return a List, but got nil instead")
	}
	_, err = d2.Get(0)
	if err == nil {
		t.Errorf("Expected Get() to return an error, but got nil instead")
	}
}

func TestGetEnumerable(t *testing.T) {
	d := From([]int{1, 2, 3})

	enumerable := d.GetEnumerable()
	assert.Equal(t, 1, enumerable.Current())
	assert.True(t, enumerable.MoveNext())
	assert.Equal(t, 2, enumerable.Current())
}
