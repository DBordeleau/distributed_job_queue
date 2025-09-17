package queue

// min-heap based on RunAt then Priority

import (
	"container/heap"
	"time"
)

type JobHeap []*Job

// Returns true if job at index i runs before job at index j,
// or if they run at the same time, if job i has higher priority.
func (h JobHeap) Less(i, j int) bool {
	if h[i].RunAt.Equal(h[j].RunAt) {
		return h[i].Priority > h[j].Priority
	}
	return h[i].RunAt.Before(h[j].RunAt)
}

// Updates RunAt/Priority and fixes the heap
func (h *JobHeap) Update(job *Job, newRunAt time.Time, newPriority int) {
	job.RunAt = newRunAt
	job.Priority = newPriority
	heap.Fix(h, job.index)
}

func (h *JobHeap) Push(x any) {
	job := x.(*Job)
	job.index = len(*h)
	*h = append(*h, job)
}

func (h *JobHeap) Pop() any {
	if h.Len() == 0 {
		return nil
	}
	old := *h
	n := len(old)
	job := old[n-1]
	*h = old[0 : n-1]
	job.index = -1
	return job
}

func (h JobHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
	h[i].index = i
	h[j].index = j
}

func (h JobHeap) Peek() *Job {
	if h.Len() == 0 {
		return nil
	}
	return h[0]
}

func (h JobHeap) Len() int {
	return len(h)
}
