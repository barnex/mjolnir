package helheim

import (
	"errors"
	"fmt"
	"io"
	"os/user"
	"strconv"
)

// Cluster user.
type User struct {
	name    string
	share   int // Relative group share of the user
	use     int // Current number of jobs running
	que     JobQueue
	group   *Group
	mailbox Mailbox
}

// API func to add new group with share.
func AddUserAPI(out io.Writer, usr *user.User, args []string) error {
	user := args[0]
	group := GetGroup(args[1])
	if group == nil {
		return errors.New("no such group: " + args[1])
	}
	share, err := strconv.Atoi(args[2])
	if err != nil {
		return err
	}
	group.AddUser(user, share)
	return nil
}

func (u *User) String() string {
	return u.name
}

func (u *User) HasJobs() bool {
	return u.que.Len() > 0
}

// Roughly the fractional use of the user.
// User with largest share wins if no jobs are running yet
func (u *User) FracUse() float64 {
	return (float64(u.use) + 1e-3) / float64(u.share)
}

// API func, prints user info.
func Users(out io.Writer) error {
	for _, u := range users {
		fmt.Fprintln(out, u)
	}
	return nil
}

func GetUser(username string) *User {
	usr, ok := users[username]
	if !ok {
		panic(errors.New("unknown username: " + username))
	}
	return usr
}
