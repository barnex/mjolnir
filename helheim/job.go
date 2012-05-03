package helheim

import (
	"fmt"
)

// Compute job.
type Job struct {
	priority int
	id       int
	file     string
}

const DEFAULT_PRIORITY = 0

var idCounter = 1

// Job constructor.
func NewJob(file string) *Job {
	job := new(Job)
	job.priority = DEFAULT_PRIORITY
	job.id = idCounter
	idCounter++
	job.file = file
	return job
}

func (j *Job) String() string {
	return fmt.Sprintf("%07d %02d  %v", j.id, j.priority, j.file)
}
