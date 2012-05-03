package helheim

import (
	"fmt"
	"io"
)

// Cluster user.
type User struct {
	name  string
	share int // Relative group share of the user
	use   int // Current number of jobs running
	que   JobQueue
	group *Group
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
