package helheim

import (
	"time"
)

const (
	SECOND       = 1e9
	HOUR         = 3600 * SECOND
	HEARTBEAT    = 60 * SECOND
	MAX_WALLTIME = 24 * HOUR
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
	//Debug("tick")
	for _, j := range running {
		if j.Walltime() > MAX_WALLTIME {
			Debug("max walltime reached for ", j)
			j.Kill()
		}
	}
}
