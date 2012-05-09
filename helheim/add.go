package helheim

import (
	"errors"
	"io"
	"os"
	"os/user"
	"path"
	"strconv"
	"strings"
)

// API func, adds job.
func Add(out io.Writer, osUser *user.User, args []string) (err error) {
	// Setup and check user
	username := osUser.Username
	usr, ok := users[username]
	if !ok {
		return errors.New("unknown username: " + username)
	}
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
	//	Debug("nice:", nice)
	//	Debug("args:", args)

	for _, arg := range args {
		// TODO: duplicate job detection using map
		file := TranslatePath(arg)
		job := NewJob(usr, file)
		job.priority = nice
		usr.que.Push(job)
	}

	FillNodes()

	return nil
}

func TranslatePath(file string) string {
	wd, err := os.Getwd()
	if err != nil {
		return file
	}
	if !path.IsAbs(file) {
		file = wd + file
		if strings.HasPrefix(file, translate[0]) {
			return translate[1] + file[len(translate[0]):]
		}
	}
	return file
}
