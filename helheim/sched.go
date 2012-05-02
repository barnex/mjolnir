package helheim

import (
	"io"
	"os/user"
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
func Add(out io.Writer, usr *user.User, args []string) error {
	Debug(usr.Username, "add", args)
	return nil
}
