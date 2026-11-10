package main

func plusOne(digits []int) []int {
	n := len(digits)

	result := NewArrayFromSlice(digits)

	for i := n - 1; i >= 0; i-- {
		if result.data[i] < 9 {
			result.data[i]++
			return result.data
		}
		result.data[i] = 0
	}

	result.PushFront(1)

	return result.data
}
