package helheim

import (
	"fmt"
	"io"
	"os/user"
)

// API func, removes job
func Rm(out io.Writer, osUser *user.User, args []string) (err error) {
	usr := GetUser(osUser.Username)
	Debug(usr.name, "rm", args)

	for _, arg := range args {
		file := TranslatePath(arg)
		ok := false

		job := usr.que.ByFilename(file)
		if job != nil {
			fmt.Fprintln(out, "rm", job)
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
	//	FillNodes()
	return nil
}
