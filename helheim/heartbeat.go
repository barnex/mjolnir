package helheim

import (
	"time"
)

const (
	SECOND                = 1e9
	MINUTE                = 60 * SECOND
	HOUR                  = 60 * MINUTE
	HEARTBEAT             = 10 * SECOND
	MAX_WALLTIME          = 24 * HOUR  // jobs get killed after running this long
	MAIL_AGGREGATION_TIME = 1 * MINUTE // aggregate mail messages for this long
)

func RunHeartbeat() {
	for {
		lock.Lock()
		heartbeat()
		lock.Unlock()
		time.Sleep(HEARTBEAT)
	}
}

func heartbeat() {
	checkWalltime()
	checkMail()
}

// Kill jobs that have been running too long.
func checkWalltime() {
	for _, j := range running {
		if j.Walltime() > MAX_WALLTIME {
			Debug("max walltime reached for ", j)
			j.Kill()
		}
	}
}

// Check if we need to send the user's aggregated mail.
func checkMail() {
	for _, usr := range users {
		// time between mails grows linearly over time, to avoid spamming.
		// mjolnir status resets time.
		if usr.mailbox.Walltime() > time.Duration(int64(usr.mailbox.nSend)*MAIL_AGGREGATION_TIME) {
			usr.mailbox.Sendmail()
		}
	}
}
