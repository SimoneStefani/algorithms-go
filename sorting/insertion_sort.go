package main

import "github.com/SimoneStefani/algorithms-go/utils"

// For each i from 1 to n-1, exchange a[i] with the entries
// that are larger in a[0] through a[i-1]. As the index i
// travels from left to right, the entries to its left are
// in sorted order in the array, so the array is fully
// sorted when i reaches the right end.
func InsertionSort(a []int) []int {
	for i := 0; i < len(a); i++ {
		for j := i; j > 0 && a[j] < a[j-1]; j-- {
			helpers.SwapInSlice(a, j, j-1)
		}
	}

	return a
}
