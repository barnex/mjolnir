package main

// This file implements the main function

import (
	"mjolnir/midgard"
	"flag"
	"fmt"
	"runtime"
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

	if *flag_version {
		fmt.Println(`Mjǫlnir 0.0.0 hard-coded cluster management`)
		fmt.Println("Go ", runtime.Version())
		return
	}

	if *flag_daemon {
		NewServer().Daemon()
		return
	}

	midgard.MainClient(flag.Args())
}
