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
	for _, gr := range groups {
		fmt.Fprint(out, gr.name, " (share ", gr.share, "/", TotalGroupShare(), ")\n")
		for _, usr := range gr.users {
			fmt.Fprint(out, "\t", usr.name, " (share ", usr.share, "/", gr.TotalUserShare(), ")\n")
		}
	}
	return nil
}
