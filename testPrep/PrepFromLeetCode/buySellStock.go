package main

//198/213 time limit exceeded

import "math"

func maxInSlice(slice []int) int {
	max := -1
	for i := 0; i < len(slice); i++ {
		max = int(math.Max(float64(max), float64(slice[i])))
	}
	return max
}

func maxProfit(prices []int) int {
	var buy int
	var sell int
	for i := 0; i < len(prices); i++ {
		val := prices[i]

		remainingMax := maxInSlice(prices[i:])

		if (remainingMax - val) > (sell - buy) {
			buy = val
			sell = remainingMax
		}
	}

	return sell - buy
}
