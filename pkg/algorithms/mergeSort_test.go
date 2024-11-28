package algorithms

import (
	"reflect"
	"testing"
)

func TestMergeSort(t *testing.T) {
	t.Run("should sort a slice of strings", func(t *testing.T) {
		strings := []string{"banana", "apple", "cherry", "date"}
		got := MergeSort(strings)
		want := []string{"apple", "banana", "cherry", "date"}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})
}
