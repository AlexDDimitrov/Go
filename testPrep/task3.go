package main

// 2. Цикличен масив 2 | 3 | 1 | 5 | 6 | 4  ... т.е на 4 ще следва 2
// 3 | 5 | 5 | 6 | -1 | 5
/// търсим първия по-голям елемент, ако няма -1

import (
	"fmt"
)

func NOTmain() {
	data := [...]int{2, 3, 1, 5, 6, 4}

	var result [6]int
	for i := 0; i < len(data); i++ {
		nextBigger := -1

		for j := 1; j < len(data); j++ {
			checkIndex := (i + j) % len(data)

			if data[checkIndex] > data[i] {
				nextBigger = data[checkIndex]
				break
			}
		}

		result[i] = nextBigger
	}
	fmt.Print(result)
}
