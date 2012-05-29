package helheim

import (
	"io"
	"os/user"
	"strings"
)

// API func, adds job.
func Add(out io.Writer, osUser *user.User, args []string) (err error) {
	usr := GetUser(osUser.Username)
	Debug(usr.name, "add", args)

	nice := 0
	args, nice, err = ParseIntFlag(args, "-pr")
	if err != nil {
		return
	}

	gpus := 0
	args, gpus, err = ParseIntFlag(args, "-gpus")
	if gpus < 1 {
		gpus = 1
	}
	if err != nil {
		return
	}

	err = CheckNoMoreFlags(args)
	if err != nil {
		return
	}

	for _, arg := range args {
		file := TranslatePath(arg)
		job := NewJob(usr, file)
		job.priority = nice
		job.gpus = gpus
		usr.que.Push(job)
	}

	FillNodes()

	return nil
}

// Translate path form headnode to compute node.
// E.g.: /diskless/home/user/file.py -> /home/user/file.py
func TranslatePath(file string) string {
	if strings.HasPrefix(file, translate[0]) {
		file = translate[1] + file[len(translate[0]):]
	}
	return file
}
