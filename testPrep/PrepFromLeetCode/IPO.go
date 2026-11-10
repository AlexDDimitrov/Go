package main

//32/40 time limit exceeded

func SortBoth(capital []int, profits []int) {
	n := len(capital)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if capital[j] > capital[j+1] {
				// Swap capital AND profits to keep them matched!
				capital[j], capital[j+1] = capital[j+1], capital[j]
				profits[j], profits[j+1] = profits[j+1], profits[j]
			}
		}
	}
}

func remove(slice []int, index int) []int {
	return append(slice[:index], slice[index+1:]...)
}

func findMaximizedCapital(k int, w int, profits []int, capital []int) int {
	SortBoth(capital, profits)

	for turn := 0; turn < k; turn++ {
		maxProfit := -1
		bestIdx := -1

		for i := 0; i < len(capital); i++ {
			if w < capital[i] {
				break
			}

			if profits[i] > maxProfit {
				maxProfit = profits[i]
				bestIdx = i
			}
		}

		if bestIdx == -1 {
			break
		}

		w += maxProfit

		capital = remove(capital, bestIdx)
		profits = remove(profits, bestIdx)
	}

	return w
}
