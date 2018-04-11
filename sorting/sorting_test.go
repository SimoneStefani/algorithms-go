package main

import (
	"testing"

	"github.com/SimoneStefani/algorithms-go/utils"
)

func TestSelectionSort(t *testing.T) {
	unordered := []int{1, 4, 2, 5, 3}
	ordered := []int{1, 2, 3, 4, 5}
	sorted := SelectionSort(unordered)

	if !helpers.AreSlicesEqual(ordered, sorted) {
		t.Error("Expected [1 2 3 4 5], got", sorted)
	}
}

func TestInsertionSort(t *testing.T) {
	unordered := []int{1, 4, 2, 5, 3}
	ordered := []int{1, 2, 3, 4, 5}
	sorted := InsertionSort(unordered)

	if !helpers.AreSlicesEqual(ordered, sorted) {
		t.Error("Expected [1 2 3 4 5], got", sorted)
	}
}
