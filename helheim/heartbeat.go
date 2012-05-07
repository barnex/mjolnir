package helheim

import (
	"time"
)

const (
	SECOND    = 1e9
	HEARTBEAT = 1 * SECOND
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
}
