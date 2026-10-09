package main

type Array[T any] struct {
	data []T
}

func NewArray[T any]() *Array[T] {
	return &Array[T]{
		data: make([]T, 0),
	}
}

func NewArrayWithCapacity[T any](capacity int) *Array[T] {
	return &Array[T]{
		data: make([]T, 0, capacity),
	}
}

func NewArrayIntWithCapacity[T any](capacity int) *Array[T] {
	return &Array[T]{
		data: make([]T, 0, capacity),
	}
}

func (a *Array[T]) IsEmpty() bool {
	return a.Len() == 0
}

func (a *Array[T]) isFull() bool {
	return a.Len() == a.Cap()
}

func (a *Array[T]) Len() int {
	return len(a.data)
}

func (a *Array[T]) Cap() int {
	return cap(a.data)
}

func (a *Array[T]) Push(value T) {
	//a.data = append(a.data, value)

	if a.isFull() {
		newCap := 2 * a.Cap()
		if newCap == 0 {
			newCap = 32
		}
		tmp := make([]T, 0, newCap)

		/*
			for _, x := range a.data {
				tmp = append(tmp, x)
			}
		*/

		/*
			for i,x := range a.data {
				tmp[i] = x
			}
		*/

		//full copy of a data
		copy(tmp, a.data)

		a.data = tmp
	}
	a.data = append(a.data, value)
}

func (a *Array[T]) Last() T {
	if a.IsEmpty() {
		panic("Array is empty")
	}
	return a.data[a.Len()-1]
}

func (a *Array[T]) First() T {
	if a.IsEmpty() {
		panic("Array is empty")
	}
	return a.data[0]
}

func (a *Array[T]) Pop() (T, bool) {
	if a.IsEmpty() {
		var zero T
		return zero, false
	}

	last := a.Last()
	a.data = a.data[:a.Len()-1]
	return last, true
}

//Stack:

type Stack[T any] struct {
	array *Array[T]
}

func NewStack[T any]() *Stack[T] {
	return &Stack[T]{
		array: NewArray[T](),
	}
}

func (s *Stack[T]) Push(value T) {
	s.array.Push(value)
}

func (s *Stack[T]) Pop() (T, bool) {
	return s.array.Pop()
}

func (s *Stack[T]) Len() int {
	return s.array.Len()
}

func (s *Stack[T]) Cap() int {
	return s.array.Cap()
}

type Queue[T any] struct {
	array *Array[T]
}

//HW to implement Queue

func main() {
	//arr := NewArray[int]()
	//arr.Push(1)
	//arr.Push(2)
	//arr.Push(3)
	//x, ok := arr.Pop()
	//if ok {
	//	println(x, ok, arr.Len(), arr.Cap()) // 3, true, 2, 32

	//if using append in Push: 3, true, 2, 4

	newStack := NewStack[int]()
	newStack.Push(10)
	newStack.Push(20)
	newStack.Push(30)
	x, ok := newStack.Pop()
	if ok {
		println(x, ok, newStack.Len(), newStack.Cap()) // 30, true, 2, 32
	}
}
