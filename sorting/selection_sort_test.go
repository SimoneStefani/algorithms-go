package main

import "testing"

func TestSelectionSort(t *testing.T) {
	unordered := []int{1, 4, 2, 5, 3}
	ordered := []int{1, 2, 3, 4, 5}
	sorted := SelectionSort(unordered)
	if !testEq(ordered, sorted) {
		t.Error("Expected [1 2 3 4 5], got", sorted)
	}
}

func testEq(a, b []int) bool {

	if a == nil && b == nil {
		return true
	}

	if a == nil || b == nil {
		return false
	}

	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}
