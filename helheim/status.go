package helheim

import (
	"fmt"
	"io"
	"os/user"
)

const (
	STATUS_QUE_LEN  = 10 // show this many entries in queue status list
	STATUS_DONE_LEN = 20 // show this many entries in done status list
)

// API func, prints job info.
func Status(out io.Writer, osUser *user.User) error {

	// Status check clears mailbox.
	// We don't want to mail what the user has already seen.
	usr := GetUser(osUser.Username)
	usr.mailbox.Clear()
	usr.mailbox.nSend = 0

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
	if donecount > 0 {
		fmt.Fprintln(out, donecount, "jobs finished:")
	}
	start := len(done) - 1
	for i := start; i >= 0; i-- {
		fmt.Fprintln(out, " ", done[i])
	}
	if donecount > STATUS_DONE_LEN {
		fmt.Fprintln(out, "...")
	}
	return nil
}
