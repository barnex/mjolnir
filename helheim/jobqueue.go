package helheim

import (
	"container/heap"
)

// Priority queue for jobs.
type JobQueue struct {
	pq          priorityQueue
	initialized bool
}

// Capacity for JobQueue.
const DEFAULT_CAP = 1000

// JobQueue constructor.
func NewJobQueue() JobQueue {
	var q JobQueue
	q.init()
	return q
}

func (jq *JobQueue) init() {
	jq.pq = make(priorityQueue, 0, DEFAULT_CAP)
	heap.Init(&(jq.pq))
	jq.initialized = true
}

// Length of the queue.
func (jq *JobQueue) Len() int {
	return jq.pq.Len()
}

// Push a new job to the queue.
func (jq *JobQueue) Push(job *Job) {
	if !jq.initialized {
		jq.init()
	}
	heap.Push(&(jq.pq), job)
}

// Pop the highest-priority job from the queue.
func (jq *JobQueue) Pop() *Job {
	if !jq.initialized {
		jq.init()
	}
	return heap.Pop(&(jq.pq)).(*Job)
}
