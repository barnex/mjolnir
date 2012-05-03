package helheim

import (
	"fmt"
)

// Compute device.
type Device struct {
	name     string
	totalMem int64
	busy     bool
	drain    bool
}

func (d *Device) String() string {
	return fmt.Sprint(d.name, " ", d.Megabytes(), "MB ", busy(d.busy))
}

func busy(b bool) string {
	if b {
		return "busy"
	}
	return "free"
}

// Total memory in megabytes.
func (d *Device) Megabytes() int {
	return int(d.totalMem / (1024 * 1024))
}
