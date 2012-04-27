package helheim

import (
	"mjolnir/midgard"
	"sync"
)

var (
	lock sync.Mutex // Protects scheduler state, pointer passed to midgard front-end
)

func MainDaemon() {

	Configure()
	initMidgard()

	go RunSched()

	// Start listening for commands
	midgard.Listen()
}

func initMidgard() {
	midgard.Lock = &lock

	midgard.Api["version"] = Version
	midgard.Help["version"] = "Print version info"

	midgard.Api["users"] = Users
	midgard.Help["users"] = "Print group and user info"
}
