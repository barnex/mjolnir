package main

// This file implements the main function

import (
	"flag"
	"mjolnir/helheim"
	"mjolnir/midgard"
)

// Command-line flags for special modes
// not normally used by the user.
var (
	flag_daemon  *bool = flag.Bool("d", false, "run in daemon mode")
	flag_version *bool = flag.Bool("v", false, "show version and exit")
)

func main() {
	flag.Parse()

	midgard.ProgName = "mjolnir"

	// Daemon mode: enter the realm of the daemons
	if *flag_daemon {
		helheim.MainDaemon()
		return
	}

	// Client mode: stay in the realm of the humans
	midgard.MainClient(flag.Args())
}
