package helheim

import (
	"fmt"
	"path"
)

// Compute job.
type Job struct {
	priority int
	id       int
	file     string
	user     *User
	node     *Node
	dev      []int
	err      error
}

const DEFAULT_PRIORITY = 0

var idCounter = 1

// Job constructor.
func NewJob(user *User, file string) *Job {
	job := new(Job)
	job.priority = DEFAULT_PRIORITY
	job.id = idCounter
	idCounter++
	job.file = file
	job.user = user
	job.node = nil
	return job
}

func (j *Job) String() string {
	str1 := fmt.Sprintf("%07d %-7s %02d  %v", j.id, j.user, j.priority, j.file)
	if j.node != nil{
		str1 += fmt.Sprint(" ", j.node, j.dev)
	}
	if j.err != nil {
		str1 += " " + j.err.Error()
	}
	return str1
}

// Working directory for job.
func(j*Job)Wd()string{
	return path.Dir(j.file)
}
