package main

// binarySearch finds the leftmost index of val if it exists.
// If val is missing, it returns the insertion index needed to maintain sorted order
func binarySearch(nums []int, val int) int {
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

func searchRange(nums []int, target int) []int {

	start := binarySearch(nums, target)
	end := binarySearch(nums, target+1) - 1

	if start < len(nums) && nums[start] == target {
		return []int{start, end}
	}

	return []int{-1, -1}

	/*found := false

	for i := 0; i < len(nums); i++ {
		if nums[i] == target {
			if !found {
				found = true
				result[0] = i
				result[1] = i
			} else {
				result[1] = i
			}
		}
	}

	if !found {
		result[0] = -1
		result[1] = -1
	}
	*/
}
