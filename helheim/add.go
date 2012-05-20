package helheim

import (
	"errors"
	"io"
	"os/user"
	"strconv"
	"strings"
)

// API func, adds job.
func Add(out io.Writer, osUser *user.User, args []string) (err error) {
	usr := GetUser(osUser.Username)
	Debug(usr.name, "add", args)

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

	for _, arg := range args {
		file := TranslatePath(arg)
		job := NewJob(usr, file)
		job.priority = nice
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
