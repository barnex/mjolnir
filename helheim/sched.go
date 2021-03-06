package helheim

import (
	"fmt"
	"io"
	"sync"
	"syscall"
	"time"
)

var (
	lock      sync.Mutex // Protects scheduler state, pointer passed to midgard front-end
	nodes     []*Node
	groups    = make(map[string]*Group)
	users     = make(map[string]*User) // Username -> User map.
	running   JobList
	done      JobList
	donecount = 0
)

// Start as many jobs as possible.
func FillNodes() {
	for DevAvailable() {
		usr := NextUser()
		if usr == nil {
			return
		} // no jobs queued 
		job := usr.que.Pop()
		node, devid := FindDevice(job)
		// no suited node/devs.
		// TODO: job with too high GPUs will block everything!
		if node == nil {
			usr.que.Push(job)
			return
		}
		Dispatch(job, node, devid)
	}
}

// Start a job on a node
func Dispatch(job *Job, node *Node, dev []int) {
	Debug("dispatch", job, "to", node, dev)

	// Bookkeeping
	job.node = node
	job.dev = dev
	job.startTime = time.Now()
	for _, d := range dev {
		node.devices[d].busy = true
	}
	job.user.use += len(dev)

	running.Append(job)

	// setup -gpu=i,j,k flag
	// TODO: make mumax-independent.
	gpuflag := fmt.Sprint("-gpu=", job.dev[0])
	for i := 1; i < len(job.dev); i++ {
		gpuflag = fmt.Sprint(gpuflag, ",", job.dev[i])
	}

	exec := executable[0]
	if job.exec != "" {
		exec = job.exec
	}
	job.cmd = job.node.Cmd("", exec, append(executable[1:], gpuflag, job.file)...)

	// Actually run the job
	go func() {
		out, err := job.cmd.CombinedOutput()
		Debug(string(out)) // TODO
		lock.Lock()
		job.err = err
		Undispatch(job)
		lock.Unlock()
	}()
}

// Handle a finished job.
func Undispatch(job *Job) {
	Debug("undispatch", job)
	defer FillNodes()

	// TODO: problematic if node has been reconfigured
	for _, d := range job.dev {
		job.node.devices[d].busy = false
	}
	job.user.use -= len(job.dev)

	running.Remove(job)

	if job.requeue {
		requeue(job)
		return
	}

	job.stopTime = time.Now()

	// Save the last few finished jobs
	donecount++
	done.Append(job)
	if len(done) > STATUS_DONE_LEN {
		done = done[len(done)-STATUS_DONE_LEN : len(done)]
	}

	// Handle failed job
	state := job.cmd.ProcessState
	if !state.Success() {
		job.user.mailbox.Println("FAIL", job)
		// On node trouble, reconfigure the node and requeue the job
		sys := state.Sys().(syscall.WaitStatus)
		if IsNodeProblem(sys.ExitStatus()) {
			job.node.Autoconf()
			requeue(job)
		}
		//TODO
	}

	if job.user.que.Len() == 0 {
		job.user.mailbox.Println("QUEUE EMPTY")
	}

}

func requeue(old *Job) {
	nw := NewJob(old.user, old.file)
	nw.exec = old.exec
	nw.cmd = old.cmd
	nw.gpus = old.gpus
	nw.id = old.id
	nw.maxWalltime = old.maxWalltime
	nw.priority = old.priority
	old.user.que.Push(nw)
}

// Reports if the job exit status signals a problem with the node itself.
// In that case, the node will need to be re-configured or rebooted.
func IsNodeProblem(exitstatus int) bool {
	return exitstatus == 255 || exitstatus == 127
}

// Find a node and GPU id(s) suited for the job.
func FindDevice(job *Job) (node *Node, dev []int) {
	dev = make([]int, job.gpus)

	for _, node = range nodes {
		found := 0
		// skip broken node
		if node.err != nil {
			continue
		}
		for i, d := range node.devices {
			if !d.busy {
				dev[found] = i
				found++
				if found == job.gpus {
					return
				}
			}
		}
	}
	return nil, nil
}

// Is there any computing device available?
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
