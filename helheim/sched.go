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

}

// Next job who gets to run
//func NextJob()*Job{
//}

func NextUser() *User {
	var nextUser *User
	leastFrac := 1e100
	for _, u := range NextGroup().users {
		if u.HasJobs() {
			if u.FracUse() < leastFrac {
				nextUser = u
				leastFrac = u.FracUse()
			}
		}
	}
	return nextUser
}
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
	running.Append(job)
	job.node = node
	for _, d := range dev {
		node.devices[d].busy = true
	}
	job.user.use++
}
