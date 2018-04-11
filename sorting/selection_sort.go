package main

// For each i, this implementation puts the (i+1)st
// smallest item in a[a]. The entries to the left of
// position i are the i smallest items in the array
// and are not examined again.
func SelectionSort(a []int) []int {
	for i := range a {
		min := i

		for j := i; j < len(a); j++ {
			if a[j] < a[min] {
				min = j
			}
		}

		Swap(a, i, min)
	}

	return a
}

// Swap elements at position i and j in an array
func Swap(a []int, i, j int) {
	k := a[i]
	a[i] = a[j]
	a[j] = k
}
