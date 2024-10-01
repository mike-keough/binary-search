package bst_test

import (
	"reflect"
	"testing"
)

func TestSort(t *testing.T) {
	got := []int{1, 7, 4, 23, 8, 9, 4, 3, 5, 7, 9, 67, 6345, 324}
	want := []int{1, 3, 4, 5, 7, 8, 9, 23, 67, 324, 6345}

	if reflect.DeepEqual(got, want) {
		t.Errorf("Got: %v Want: %v", got, want)
	}
}

func TestBuildTree(t *testing.T) {

}
