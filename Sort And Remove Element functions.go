package main

func remove(slice []int, index int) []int {
	return append(slice[:index], slice[index+1:]...)
}

// Normal sorts
func Sort1D(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			// Compare adjacent elements directly
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}

// Sorts all values
func Sort2D(grid [][]int) {
	// 2D - 1D
	var flat []int
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			flat = append(flat, grid[i][j])
		}
	}

	Sort1D(flat)

	// 1D - 2D
	index := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			grid[i][j] = flat[index]
			index++
		}
	}
}

// Sorts by the first element of the sub array
func Sort2DFE(intervals [][]int) {
	n := len(intervals)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			// Compare adjacent elements using their first element index [0]
			if intervals[j][0] > intervals[j+1][0] {
				intervals[j], intervals[j+1] = intervals[j+1], intervals[j]
			}
		}
	}
}
