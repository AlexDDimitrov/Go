package main

import "fmt"

type RingBuffer[T any] struct {
	data  []T
	first int
	last  int
	size  int
}

func NewRingBuffer[T any](capacity int) *RingBuffer[T] {
	if capacity < 1 {
		panic(fmt.Sprintf("expecting capacity to be positive integer, got: %d", capacity))
	}
	return &RingBuffer[T]{
		data:  make([]T, capacity),
		first: 0,
		last:  0,
		size:  0,
	}
}

func (rb *RingBuffer[T]) IsEmpty() bool {
	return rb.size == 0
}

func (rb *RingBuffer[T]) IsFull() bool {
	return rb.size == cap(rb.data)
}

func (rb *RingBuffer[T]) Enqueue(x T) {
	if rb.IsFull() {
		newCapacity := 2 * cap(rb.data)
		tmp := make([]T, newCapacity)
		copy(tmp, rb.data[rb.first:cap(rb.data)])
		if rb.first != 0 {
			copy(tmp[rb.first+1:], rb.data[0:rb.first+1])
		}
		rb.last = cap(rb.data)
		rb.data = tmp
		rb.first = 0
	}
	rb.data[rb.last] = x
	rb.last = (rb.last + 1) % cap(rb.data)
	rb.size++
}

func (rb *RingBuffer[T]) Dequeue() (T, bool) {
	if rb.IsEmpty() {
		var tmp T
		return tmp, false
	}
	x := rb.data[rb.first]
	rb.first = (rb.first + 1) % cap(rb.data)
	rb.size--
	return x, true
}

func (rb *RingBuffer[T]) First() (T, bool) {
	if rb.IsEmpty() {
		var zero T
		return zero, false
	}
	return rb.data[rb.first], true
}

func (rb *RingBuffer[T]) Last() (T, bool) {
	if rb.IsEmpty() {
		var zero T
		return zero, false
	}
	return rb.data[(rb.last-1)%cap(rb.data)], true
}

func (rb *RingBuffer[T]) Data() []T {
	res := make([]T, rb.size)

	for i := 0; i < rb.size; i++ {
		res[i] = rb.data[(rb.first+i)%cap(rb.data)]
	}

	return res
}

func maxSubarraySumCircular(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	// Step 1: Put all the numbers into your RingBuffer
	rb := NewRingBuffer[int](len(nums))
	for _, x := range nums {
		rb.Enqueue(x)
	}

	// Step 2: Unroll the circular data back into a straight line
	// This lets us read the elements in their true logical order
	orderedNums := rb.Data()

	// Step 3: Run the calculations
	totalSum := 0
	maxSum, minSum := orderedNums[0], orderedNums[0]
	currentMax, currentMin := 0, 0

	for _, x := range orderedNums {
		totalSum += x

		// Case 1: Track the highest standard subarray sum (no wrapping)
		currentMax = max(x, currentMax+x)
		maxSum = max(maxSum, currentMax)

		// Case 2: Track the lowest middle subarray sum (for wrapping)
		currentMin = min(x, currentMin+x)
		minSum = min(minSum, currentMin)
	}

	// Step 4: Handle the all-negative edge case
	// If the highest sum is negative, the whole array is negative.
	if maxSum < 0 {
		return maxSum
	}

	// Step 5: Choose between the straight line max vs. the wrapped max
	wrappedMax := totalSum - minSum
	return max(maxSum, wrappedMax)
}
