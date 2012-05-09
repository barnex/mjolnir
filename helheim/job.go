package helheim

import (
	"fmt"
	"os/exec"
	"path"
	"time"
)

// Compute job.
type Job struct {
	priority  int       // User-defined job priority
	id        int       // Unique job id
	index     int       // Index in the priority queue (internal use)
	file      string    // Mumax input file
	user      *User     // User who owns job
	node      *Node     // Node assigned to job, if any yet
	dev       []int     // GPU device indices assigned to job, if any yet
	err       error     // Error executing job, if any
	cmd       *exec.Cmd // Command executing the job
	startTime time.Time // Walltime when job was started
	stopTime  time.Time // Walltime when job was stoped
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
	wall := j.Walltime()
	str1 := fmt.Sprintf("%02d %07d %-7s %02d %v %v", j.index, j.id, j.user, j.priority, formatDuration(wall), j.file)
	if j.node != nil {
		str1 += fmt.Sprint(" ", j.node, j.dev)
	}
	if j.err != nil {
		str1 += " " + j.err.Error()
	}
	return str1
}

func formatDuration(wall time.Duration) string {
	return fmt.Sprintf("%02d:%02d:%02d", int(wall.Hours()), int(wall.Minutes())%60, int(wall.Seconds())%60)
}

// Working directory for job.
func (j *Job) Wd() string {
	return path.Dir(j.file)
}

func (j *Job) Running() bool {
	return !j.startTime.IsZero() && j.stopTime.IsZero()
}

func (j *Job) Walltime() time.Duration {
	if j.Running() {
		return time.Now().Sub(j.startTime)
	}
	return j.stopTime.Sub(j.startTime)
}
