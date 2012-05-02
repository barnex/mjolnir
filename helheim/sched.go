package helheim

import (
	"io"
)

var (
	nodes  []*Node
	groups []*Group
	users  = make(map[string]*User) // Username -> User map.
)

// Run the scheduler. Infinite loop.
func RunSched() {
	/*FillNodes()
	for {
		select {
		case done := <-finish:
			Undispatch(done.Job, done.exitStatus)
			FillNodes()
		}
	}*/
}

// API func, adds job.
func Add(out io.Writer, args []string) error {
	Debug("add", args)
	return nil
}
