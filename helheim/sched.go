package helheim

import (
	"sync"
)

var (
	Lock sync.Mutex // Protects scheduler state, pointer passed to midgard front-end
)
