package helheim

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

//func NextJob()*Job{
//	
//}

func Dispatch(job *Job, node *Node, dev []int) {
	running.Append(job)
	job.node = node
	for _, d := range dev {
		node.devices[d].busy = true
	}
}
