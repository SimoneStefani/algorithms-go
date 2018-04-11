package helpers

// Swap elements at position i and j in a slice
func SwapInSlice(a []int, i, j int) {
	k := a[i]
	a[i] = a[j]
	a[j] = k
}

// Check if two slices are equal
func AreSlicesEqual(a, b []int) bool {
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
