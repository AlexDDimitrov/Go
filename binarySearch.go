package main

// standard binarySearch
func binarySearch(nums []int, target int) int {
	low := 0
	high := len(nums) - 1

	for low <= high {
		mid := low + (high-low)/2

		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	return -1
}

// binarySearch finds the leftmost index of val if it exists.
// If val is missing, it returns the insertion index needed to maintain sorted order
func binarySearchIndex(nums []int, val int) int {
	left, right := 0, len(nums)
	for left < right {
		mid := left + (right-left)/2
		if nums[mid] < val {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return left
}
