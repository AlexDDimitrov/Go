package main

func Sort1D(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			// Compare adjacent elements directly
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}

func maxTaskAssign(tasks []int, workers []int, pills int,
	strength int) int {
	Sort1D(tasks)
	Sort1D(workers)

	left := 0
	right := min(len(tasks), len(workers))

	for left < right {
		mid := left + (right-left+1)/2

		if canAssign(tasks, workers, pills, strength, mid) {
			left = mid
		} else {
			right = mid - 1
		}
	}

	return left
}

func canAssign(tasks []int, workers []int, pills int, strength int, tasksCount int) bool {
	rb := NewRingBuffer[int](tasksCount)

	for i := 0; i < tasksCount; i++ {
		workerIndex := findWorker(workers, tasks[i])
		if workerIndex == -1 {
			return false
		}

		rb.Enqueue(workerIndex)

		if rb.size > 1 && workers[rb.first]+strength < tasks[i] {
			pills--
			if pills < 0 {
				return false
			}
			rb.Dequeue()
		}

		if rb.IsFull() {
			rb.Dequeue()
		}
	}

	return true
}

func findWorker(workers []int, task int) int {
	for i := 0; i < len(workers); i++ {
		if workers[i] >= task {
			return i
		}
	}
	return -1
}
