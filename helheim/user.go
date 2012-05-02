package helheim

import (
	"fmt"
	"io"
)

// Cluster user.
type User struct {
	name  string
	share int // Relative group share of the user
	que   JobQueue
}

// API func, prints user info.
func Users(out io.Writer) error {
	for _, u := range users {
		fmt.Fprintln(out, u)
	}
	return nil
}
