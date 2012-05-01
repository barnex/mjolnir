package helheim

import (
	"container/heap"
)

type JobQueue struct {
	pq          PriorityQueue
	initialized bool
}

const DEFAULT_CAP = 1000

func (jq *JobQueue) init() {
	jq.pq = make(PriorityQueue, 0, DEFAULT_CAP)
	heap.Init(&(jq.pq))
	jq.initialized = true
}

func (jq *JobQueue) Len() int {
	return jq.pq.Len()
}

func (jq *JobQueue) Push(job *Job) {
	if !jq.initialized {
		jq.init()
	}
	heap.Push(&(jq.pq), job)
}

func (jq *JobQueue) Pop() *Job {
	if !jq.initialized {
		jq.init()
	}
	return heap.Pop(&(jq.pq)).(*Job)
}
