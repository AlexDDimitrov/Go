package Stack

import Array "github.com/NoHomey/algorithms-and-data-structures-2026/dsa/Array"

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
