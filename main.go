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

const PROG = "mjolnir"

func main() {
	flag.Parse()

	if *flag_daemon {
		midgard.Api["version"] = helheim.Version
		midgard.MainDaemon()
		return
	}

	midgard.MainClient(flag.Args())
}
