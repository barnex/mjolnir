package helheim

import (
	"fmt"
	"io"
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

func (j*Job)String()string{
	return fmt.Sprintf("%07d %02d  %v", j.id, j.priority, j.file)
}

// API func, prints job info.
func Status(out io.Writer) error {
	for _, usr := range users {
		if usr.que.Len() == 0 {
			continue
		}
		fmt.Fprintln(out, usr)
		fmt.Fprintln(out, "  ID      PR  FILE")
		for _, job := range usr.que.pq {
			fmt.Fprintln(out, " ", job)
		}
	}
	return nil
}
