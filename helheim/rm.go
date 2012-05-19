package helheim

import (
	"errors"
	"fmt"
	"io"
	"os/user"
)

// API func, adds job.
func Rm(out io.Writer, osUser *user.User, args []string) (err error) {
	// Setup and check user
	username := osUser.Username
	usr, ok := users[username]
	if !ok {
		return errors.New("unknown username: " + username)
	}
	Debug(usr.name, "rm", args)

	for _, arg := range args {
		file := TranslatePath(arg)
		job := usr.que.ByFilename(file)
		ok := false
		if job != nil {
			usr.que.Remove(job)
			ok = true
		}
		for i, r := range running {
			if r.file == file {
				fmt.Fprintln(out, "kill", r)
				running[i].Kill()
				ok = true
				break
			}
		}
		if !ok {
			fmt.Fprintln(out, "no such job:", file)
		}
	}
	return nil
}
