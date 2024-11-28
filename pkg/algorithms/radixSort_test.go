package algorithms

import (
	"reflect"
	"testing"
)

func TestRadixSort(t *testing.T) {
	t.Run("should sort a slice of strings", func(t *testing.T) {
		strings := []string{"banana", "apple", "cherry", "date"}
		got := RadixSort(strings)
		want := []string{"apple", "banana", "cherry", "date"}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})
}
