package midgard

import (
	"sync"
)

var (
	ProgName    = ""                           // The program name.
	Port        = ":2728"                      // Default RPC port.
	Api         = make(map[string]interface{}) // List of available functions to user.
	Help        = make(map[string]string)      // Store help for commands here.
	Lock        *sync.Mutex                    // Protects callee state.
	SpawnDaemon = false                        // Spawn daemon if not yet running?
)
