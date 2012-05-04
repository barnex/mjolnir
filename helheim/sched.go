package helheim

import (
	"fmt"
	"io"
)

var (
	nodes   []*Node
	groups  []*Group
	users   = make(map[string]*User) // Username -> User map.
	running JobList
	done    JobList
)

// Run the scheduler. Infinite loop.
func RunSched() {
	FillNodes()
	/*for {
		select {
		case done := <-finish:
			Undispatch(done.Job, done.exitStatus)
			FillNodes()
		}
	}*/
}

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

func PrintNext(out io.Writer) error {
	fmt.Fprintln(out, "next group:", NextGroup())
	fmt.Fprintln(out, "next user:", NextUser())
	return nil
}

func Dispatch(job *Job, node *Node, dev []int) {
	Debug("dispatch", job, "to", node, dev)
	running.Append(job)
	job.node = node
	for _, d := range dev {
		node.devices[d].busy = true
	}
	job.user.use++
}
