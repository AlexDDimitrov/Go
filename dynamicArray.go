package main

type ArrayInt struct {
	data []int
}

func NewArrayInt() *ArrayInt {
	return &ArrayInt{
		data: make([]int, 0),
	}
}

func NewArrayIntWithCapacity(capacity int) *ArrayInt {
	return &ArrayInt{
		data: make([]int, 0, capacity),
	}
}

func (a *ArrayInt) IsEmpty() bool {
	return len(a.data) == 0
}

func (a *ArrayInt) Len() int {
	return len(a.data)
}

func (a *ArrayInt) Cap() int {
	return cap(a.data)
}

func (a *ArrayInt) Push(value int) {
	a.data = append(a.data, value)
}
