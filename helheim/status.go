package helheim

import (
	"fmt"
	"io"
)

const (
	STATUS_QUE_LEN  = 10 // show this many entries in queue status list
	STATUS_DONE_LEN = 10 // show this many entries in done status list
)

// API func, prints job info.
func Status(out io.Writer) error {
	// running
	fmt.Fprintln(out, len(running), "jobs running:")
	for _, job := range running {
		fmt.Fprintln(out, " ", job)
	}

	// queued
	for _, usr := range users {
		if usr.que.Len() == 0 {
			continue
		}
		fmt.Fprintln(out, usr.que.Len(), "jobs queued for", usr, ":")
		//fmt.Fprintln(out, "  ID      USER    PR  TIME      FILE")
		for i, job := range usr.que.pq {
			fmt.Fprintln(out, " ", job)
			if i == STATUS_QUE_LEN {
				fmt.Fprintln(out, "...")
				break
			}
		}
	}

	// done
	if len(done) > 0 {
		fmt.Fprintln(out, len(done), "jobs finished:")
	}
	for i, job := range done {
		fmt.Fprintln(out, " ", job)
		if i == STATUS_DONE_LEN {
			fmt.Fprintln(out, "...")
			break
		}
	}
	return nil
}
