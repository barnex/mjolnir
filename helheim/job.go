package helheim

import (
	"io"
)

// Compute job.
type Job struct {
	priority int
	id       int
	username string
	file     string
}

const DEFAULT_PRIORITY = 0

var idCounter = 1

// Job constructor.
func NewJob(username string, file string) *Job {
	job := new(Job)
	job.priority = DEFAULT_PRIORITY
	job.id = idCounter
	idCounter++
	job.username = username
	job.file = file
	return job
}

//func (j*Job)String()string{
//	return fmt.Sprint()
//}

// API func, prints job info.
func Status(out io.Writer) error {
	return nil
}
