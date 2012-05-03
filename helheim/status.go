package helheim

import (
	"fmt"
	"io"
)

// API func, prints job info.
func Status(out io.Writer) error {
	for _, usr := range users {
		if usr.que.Len() == 0 {
			continue
		}
		fmt.Fprintln(out, usr)
		fmt.Fprintln(out, "  ID      PR  FILE")
		for _, job := range usr.que.pq {
			fmt.Fprintln(out, " ", job)
		}
	}
	return nil
}
