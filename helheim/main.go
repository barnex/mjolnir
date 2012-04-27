package helheim

import (
	"mjolnir/midgard"
	"sync"
)

var (
	Lock sync.Mutex // Protects scheduler state, pointer passed to midgard front-end
)

func MainDaemon() {
	midgard.Lock = &Lock

	midgard.Api["version"] = Version
	midgard.Help["version"] = "Print version info"

	midgard.MainDaemon()
}
