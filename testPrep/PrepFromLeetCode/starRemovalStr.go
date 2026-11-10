package main

type Stack[T any] struct {
	Array *Array[T]
}

func NewStack[T any]() *Stack[T] {
	return &Stack[T]{Array: NewArray[T]()}
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

func removeStars(s string) string {
	result := NewStack[byte]()
	for i := 0; i < len(s); i++ {
		if s[i] == '*' {
			result.Pop()
		} else {
			result.Push(s[i])
		}
	}
	return string(result.Array.data)
}
