package helheim

import (
	"fmt"
	"io"
	"os/user"
)

// API func, kills job and adds it to the queue again.
func Requeue(out io.Writer, osUser *user.User, args []string) (err error) {
	// Setup and check user
	usr := GetUser(osUser.Username)
	Debug(usr.name, "requeue", args)

	for _, arg := range args {
		file := TranslatePath(arg)
		ok := false
		for i, r := range running {
			if r.file == file {
				Debug("requeue", r)
				fmt.Fprintln(out, "requeue", r)
				running[i].Requeue()
				ok = true
				break
			}
		}
		if !ok {
			fmt.Fprintln(out, "no such job:", file)
		}
	}
	//FillNodes()
	return nil
}
