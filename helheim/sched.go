package helheim

import (
	"fmt"
	"io"
	"sync"
	//"time"
)

var (
	lock    sync.Mutex // Protects scheduler state, pointer passed to midgard front-end
	nodes   []*Node
	groups  []*Group
	users   = make(map[string]*User) // Username -> User map.
	running JobList
	done    JobList
)

const (
	SECOND    = 1e9
	HEARTBEAT = 1 * SECOND
)

// Run the scheduler. Infinite loop.
//func RunSched() {
//	for {
//		lock.Lock()
//		FillNodes()
//		lock.Unlock()
//		time.Sleep(HEARTBEAT)
//	}
//}

// Start as many jobs as possible.
func FillNodes() {
	for DevAvailable() {
		usr := NextUser()
		if usr == nil {
			return
		} // no jobs queued 
		job := usr.que.Pop()
		node, devid := FindDevice(job)
		Dispatch(job, node, devid)
	}
}

// Start a job on a node
func Dispatch(job *Job, node *Node, dev []int) {
	Debug("dispatch", job, "to", node, dev)

	// Bookkeeping
	job.node = node
	job.dev = dev
	for _, d := range dev {
		node.devices[d].busy = true
	}
	job.user.use += len(dev)

	running.Append(job)

	// Actually run the job
	go Exec(job)
}

func Exec(job *Job) {
	_, err := job.node.Exec(MUMAX2, job.file)

	lock.Lock()

	job.err = err
	Undispatch(job)

	lock.Unlock()
}

func Undispatch(job *Job) {
	Debug("undispatch", job)
	for _, d := range job.dev {
		job.node.devices[d].busy = false
	}
	job.user.use -= len(job.dev)

	running.Remove(job)
	done.Append(job)
	FillNodes()
}

// Find a device and GPU id(s) suited for the job.
func FindDevice(job *Job) (node *Node, dev []int) {
	for _, n := range nodes {
		for i, d := range n.devices {
			if !d.busy {
				return n, []int{i}
			}
		}
	}
	return nil, nil
}

// Is a computing device available?
func DevAvailable() bool {
	for _, n := range nodes {
		for _, d := range n.devices {
			if !d.busy {
				return true
			}
		}
	}
	return false
}

// Next user who gets to run a job.
func NextUser() *User {
	nextGroup := NextGroup()
	if nextGroup == nil {
		return nil
	}
	var nextUser *User
	leastFrac := 1e100
	for _, u := range nextGroup.users {
		if u.HasJobs() {
			if u.FracUse() < leastFrac {
				nextUser = u
				leastFrac = u.FracUse()
			}
		}
	}
	return nextUser
}

// Next group who gets to run a job.
func NextGroup() *Group {
	var nextGroup *Group
	leastFrac := 1e100
	for _, g := range groups {
		if g.HasJobs() {
			if g.FracUse() < leastFrac {
				nextGroup = g
				leastFrac = g.FracUse()
			}
		}
	}
	return nextGroup
}

// API func, prints the next one who gets to run a job.
func PrintNext(out io.Writer) error {
	fmt.Fprintln(out, "next group:", NextGroup())
	fmt.Fprintln(out, "next user:", NextUser())
	return nil
}
