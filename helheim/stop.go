package helheim

import (
	"fmt"
	"io"
	"os/user"
)

func Please(out io.Writer, osUser *user.User, args []string) (err error) {

	// Setup and check user
	usr := GetUser(osUser.Username)
	Debug(usr.name, "requeue", args)

	cmd := args[0]
	switch cmd {
	case "stop":
		err = stop(out)
	default:
		err = NewError("yes please?")
	}
	return
}

func stop(out io.Writer) (err error) {
	for i := range nodes {
		nodes[i].err = NewError("draining")
	}
	for i := range running {
		fmt.Println("requeue", running[i])
		running[i].Requeue()
	}
	return
}
