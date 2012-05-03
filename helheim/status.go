package helheim

import (
	"fmt"
	"io"
)

// API func, prints job info.
func Status(out io.Writer) error {
	if len(running) > 0 {
		fmt.Fprintln(out, "running:")
	}
	for _, job := range running {
		fmt.Fprintln(out, " ", job)
	}

	for _, usr := range users {
		if usr.que.Len() == 0 {
			continue
		}
		fmt.Fprintln(out, "queue for", usr, ":")
		fmt.Fprintln(out, "  ID      USER    PR  FILE")
		for _, job := range usr.que.pq {
			fmt.Fprintln(out, " ", job)
		}
	}
	return nil
}
