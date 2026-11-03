package main

// 1. Списък от цели числа -> 2 | -1 | 3 | 2 | 5 | 7 | 6 | 8 | 11
// K = 4 (1<=k<len(arr))
// k = 4, len = 9
// 3 | 3 | 7 | 7 | 8 | 11

import (
	"fmt"
	"math"

	"github.com/NoHomey/algorithms-and-data-structures-2026/dsa/Array"
)

type Stack[T any] struct {
	Array *Array.Array[T]
}

func NewStack[T any]() *Stack[T] {
	return &Stack[T]{Array: Array.NewArray[T]()}
}

func (s Stack[T]) IsEmpty() bool {
	return s.Array.IsEmpty()
}

func (s Stack[T]) IsFull() bool {
	return s.Array.IsFull()
}

func (s Stack[T]) Push(x T) {
	s.Array.Push(x)
}

func (s Stack[T]) Pop() (T, bool) {
	return s.Array.Pop()
}

func (s Stack[T]) Top() T {
	return s.Array.Last()
}

func (s Stack[T]) Len() int {
	return s.Array.Len()
}

func (s Stack[T]) Cap() int {
	return s.Array.Cap()
}

func Notmain() {
	k := 4
	l := 9

	arrayData := [...]int{2, -1, 3, 2, 5, 7, 6, 8, 11}

	stack := NewStack[int]()
	for _, val := range arrayData {
		stack.Push(val)
	}

	var result [6]int

	for i := 0; i <= l-k; i++ {
		var flag bool

		data := [4]int{}

		for j := 0; j < k; j++ {
			data[j], flag = stack.Array.Get(uint(i + j))
			if !flag {
				panic("Error")
			}
		}

		check1 := math.Max(float64(data[0]), float64(data[1]))
		check2 := math.Max(float64(data[2]), float64(data[3]))
		result[i] = int(math.Max(check1, check2))
	}

	fmt.Print(result)
}
