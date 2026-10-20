package main

/*
//Queue
type Queue[T any] struct {
	array *Array[T]
	head  int
	tail  int
	size  int
}

func (q *Queue[T]) String() string {
	return fmt.Sprintf("Queue{array: %v, head: %d, tail: %d, size: %d}", q.array.data, q.head, q.tail, q.size)
}

func NewQueue[T any](capacity int) *Queue[T] {
	if capacity <= 0 {
		capacity = 8
	}
	arr := NewArrayWithCapacity[T](capacity)
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

func (q *Queue[T]) Enqueue(value T) bool {
	if q.IsFull() {
		newCap := 2 * q.array.Cap()
		if newCap == 0 {
			newCap = 8
		}

		newData := make([]T, newCap)

		for i := 0; i < q.size; i++ {
			oldIndex := (q.head + i) % q.array.Cap()
			newData[i] = q.array.data[oldIndex]
		}
		q.array.data = newData
		q.head = 0
		q.tail = q.size
	}

	q.array.data[q.tail] = value
	q.tail = (q.tail + 1) % q.array.Cap()
	q.size++
	return true
}

func (q *Queue[T]) Dequeue() (T, bool) {
	if q.IsEmpty() {
		var zero T
		return zero, false
	}

	value := q.array.data[q.head]
	var zero T
	q.array.data[q.head] = zero

	q.head = (q.head + 1) % q.array.Cap()
	q.size--
	return value, true
}
*/
