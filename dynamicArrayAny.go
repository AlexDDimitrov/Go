package main

import "fmt"

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

/*type Queue[T any] struct {
	array *Array[T]
}*/

//HW to implement Queue

type Queue[T any] struct {
	array *Array[T]
	head  int // Index of the front element
	tail  int // Index of the next insertion point (rear)
	size  int // Current number of elements in the queue
}

func (q *Queue[T]) String() string {
	return fmt.Sprintf("Queue{array: %v, head: %d, tail: %d, size: %d}", q.array.data, q.head, q.tail, q.size)
}

// NewQueue initializes a circular queue with a fixed capacity
func NewQueue[T any](capacity int) *Queue[T] {
	if capacity <= 0 {
		capacity = 8 // Default fallback capacity
	}
	arr := NewArrayWithCapacity[T](capacity)
	// Resize data length to match capacity so we can use direct index assignment
	arr.data = make([]T, capacity)

	return &Queue[T]{
		array: arr,
		head:  0,
		tail:  0,
		size:  0,
	}
}

func (q *Queue[T]) IsEmpty() bool {
	return q.size == 0
}

func (q *Queue[T]) IsFull() bool {
	return q.size == q.array.Cap()
}

func (q *Queue[T]) Len() int {
	return q.size
}

// Enqueue adds an element to the rear of the queue
func (q *Queue[T]) Enqueue(value T) bool {
	if q.IsFull() {
		newCap := 2 * q.array.Cap()
		if newCap == 0 {
			newCap = 8 // Default fallback capacity
		}

		// Create a brand new slice with the expanded capacity
		newData := make([]T, newCap)

		// Unroll the circular elements into the new slice in straight sequential order
		for i := 0; i < q.size; i++ {
			// Find old index using modulo
			oldIndex := (q.head + i) % q.array.Cap()
			newData[i] = q.array.data[oldIndex]
		}

		// Reset your tracking pointers
		q.array.data = newData
		q.head = 0
		q.tail = q.size // The next insert position is right after the old elements
	}

	q.array.data[q.tail] = value
	// Move tail forward circularly: tail = (tail + 1) % capacity
	q.tail = (q.tail + 1) % q.array.Cap()
	q.size++
	return true
}

// Dequeue removes and returns an element from the front of the queue
func (q *Queue[T]) Dequeue() (T, bool) {
	if q.IsEmpty() {
		var zero T
		return zero, false // Queue underflow
	}

	value := q.array.data[q.head]
	// Clear slot reference for garbage collection (especially useful for pointers)
	var zero T
	q.array.data[q.head] = zero

	// Move head forward circularly: head = (head + 1) % capacity
	q.head = (q.head + 1) % q.array.Cap()
	q.size--
	return value, true
}

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

	fmt.Println("Testing queue")
	newQueue := NewQueue[int](4)
	newQueue.Enqueue(10)
	newQueue.Enqueue(20)
	newQueue.Enqueue(30)
	fmt.Println(newQueue) // Queue{array: [10 20 30 0], head: 0, tail: 3, size: 3}
	y, ok := newQueue.Dequeue()
	fmt.Println(newQueue)                                                // Queue{array: [0 20 30 0], head: 1, tail: 3, size: 2}
	fmt.Println("Starting element:", newQueue.array.data[newQueue.head]) // 20
	if ok {
		fmt.Println(y, ok, newQueue.Len(), newQueue.IsFull()) // 10, true, 2, false
	}
	newQueue.Enqueue(40)
	fmt.Println(newQueue) // Queue{array: [0 20 30 40], head: 1, tail: 0, size: 3}
	newQueue.Enqueue(50)
	fmt.Println(newQueue) // Queue{array: [50 20 30 40], head: 1, tail: 1, size: 4}
	newQueue.Dequeue()
	fmt.Println(newQueue) // Queue{array: [50 0 30 40], head: 2, tail: 1, size: 3}
	newQueue.Enqueue(60)
	fmt.Println(newQueue) // Queue{array: [30 40 50 60], head: 2, tail: 2, size: 4}
	newQueue.Enqueue(70)  // This should trigger a resize
	fmt.Println(newQueue) // Queue{array: [30 40 50 60 70 0 0 0], head: 0, tail: 5, size: 5}

}
