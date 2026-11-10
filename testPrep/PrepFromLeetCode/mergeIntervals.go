package main

import "sort"

func merge(intervals [][]int) [][]int {
	if len(intervals) == 0 {
		return [][]int{}
	}

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	// 1. Instantiate using [[]int] so it accepts slice intervals
	merged := NewArrayWithCap[[]int](len(intervals))

	// 2. Process each interval
	for _, current := range intervals {
		// If merged is empty or there is no overlap, push the current interval
		if merged.IsEmpty() || current[0] > merged.Last()[1] {
			merged.Push(current)
		} else {
			// There is an overlap, update the end time of the last interval
			if current[1] > merged.Last()[1] {
				merged.data[merged.Len()-1][1] = current[1]
			}
		}
	}

	return merged.data
}
