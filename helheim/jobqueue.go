package helheim

import (
	"container/heap"
)

// Priority queue for jobs.
type JobQueue struct {
	pq          priorityQueue
	byname      map[string]*Job // maps filename onto job
	byidx       map[int]*Job    // maps id onto job
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
	jq.byname = make(map[string]*Job)
	jq.byidx = make(map[int]*Job)
	jq.initialized = true
}

// Find a job by its filename
func (jq *JobQueue) ByFilename(file string) *Job {
	if !jq.initialized {
		jq.init()
	}
	return jq.byname[file]
}

// Find a job by its ID
func (jq *JobQueue) ByID(index int) *Job {
	if !jq.initialized {
		jq.init()
	}
	return jq.byidx[index]
}

// Length of the queue.
func (jq *JobQueue) Len() int {
	return jq.pq.Len()
}

// Push a new job to the queue.
// It is safe to push jobs with file names that are already in the queue.
// They will not be added twice but the pre-existing job will be overwritten
// (apart form its ID and index in the queue)
func (jq *JobQueue) Push(job *Job) {
	if !jq.initialized {
		jq.init()
	}
	// Job with same name already present?
	if prev, ok := jq.byname[job.file]; ok {
		previndex := prev.index
		previd := prev.id
		*prev = *job           // overwrite existing job contents with new job
		prev.index = previndex // but preserve index in priority queue!
		prev.id = previd       // but preserve id
	} else {
		heap.Push(&(jq.pq), job)  // push to heap
		jq.byname[job.file] = job // and map
		jq.byidx[job.id] = job    // and map
	}
}

// Pop the highest-priority job from the queue.
func (jq *JobQueue) Pop() *Job {
	if !jq.initialized {
		jq.init()
	}
	job := heap.Pop(&(jq.pq)).(*Job)
	delete(jq.byname, job.file)
	delete(jq.byidx, job.id)
	return job
}

// Remove a job from the queue.
func (jq *JobQueue) Remove(job *Job) {
	Debug("rm", job)
	delete(jq.byname, job.file)
	delete(jq.byidx, job.id)
	heap.Remove(&(jq.pq), job.index)
	job.index = -1
}
