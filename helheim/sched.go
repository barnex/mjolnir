package helheim

import (
	"errors"
	"io"
	"os/user"
	"strconv"
	"strings"
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
func Add(out io.Writer, osUser *user.User, args []string) (err error) {
	// Setup and check user
	username := osUser.Username
	usr, ok := users[username]
	Debug(usr.name, "add", args)
	if !ok {
		return errors.New("unknown username: " + username)
	}

	// Parse "-pr" flag
	nice := DEFAULT_PRIORITY
	nicei := -1
	for i, arg := range args {
		if strings.HasPrefix(arg, "-") {
			if arg == "-pr" {
				if i == len(args)-1 {
					return errors.New("-pr needs argument")
				}
				nice, err = strconv.Atoi(args[i+1])
				nicei = i
				if err != nil {
					return
				}
				break
			} else {
				return errors.New("unknown option: " + arg + ". usage: add -pr <N> file")
			}
		}
	}
	// remove flag from args list
	if nicei != -1 {
		args = append(args[:nicei], args[nicei+2:]...)
	}
//	Debug("nice:", nice)
//	Debug("args:", args)

	for _, arg := range args {
		// TODO: duplicate job detection using map
		job := NewJob(arg)
		job.priority = nice
		usr.que.Push(job)
	}

	return nil
}
