// За масив от 10 числа искаме да намерим индексите на подмасиви, които са поредни числя
// [2, 1, 2, 4, 5, 6, 2, 3, 5]
// масивите от поредни числа [2],  [1, 2],  [4, 5, 6] [2, 3],  [5]
// -> [0:0] [1,2] [3,5] [6,7] [8]

// go mod init exec
// go.mod:
// module github.com/NoHomey/algorithms-and-data-structures-2026/dsa
//
// go 1.26.4

// terminal:
// go get github.com/NoHomey/algorithms-and-data-structures-2026/dsa
// -> to access remote libs

// and we import: Queue "github.com/NoHomey/algorithms-and-data-structures-2026/dsa/ringbuffer"
package main

import (
	"fmt"

	Queue "github.com/NoHomey/algorithms-and-data-structures-2026/dsa/ringbuffer"
)

type Range struct {
	Beg int
	End int
}

func solve(arr []int) []Range {
	q := Queue.NewRingBuffer[*Range](8)
	for i, x := range arr {
		if q.IsEmpty() {
			q.Enqueue(&Range{Beg: i, End: i})
			continue
		}
		last, ok := q.Last()
		if !ok {
			panic("Error")
		}
		if arr[last.End]+1 == x {
			last.End = i
		} else {
			q.Enqueue(&Range{Beg: i, End: i})
		}
	}

	data := q.Data()
	res := make([]Range, len(data))
	for i := range data {
		res[i] = *data[i]
	}
	return res
}

func main() {
	arr := []int{2, 1, 2, 4, 5, 6, 2, 3, 5}
	res := solve(arr)
	fmt.Println(arr)
	fmt.Println(res)
}
