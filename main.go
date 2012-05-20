package main

// This file implements the main function

import (
	"flag"
	"mjolnir/helheim"
	"mjolnir/midgard"
	"os"
	"path"
)

// Command-line flags for special modes
// not normally used by the user.
var (
	flag_daemon  = flag.Bool("d", false, "run in daemon mode")
	flag_version = flag.Bool("v", false, "show version and exit")
	flag_port    = flag.String("p", ":2527", "TCP port")
)

func main() {
	flag.Parse()

	midgard.ProgName = "mjolnir"
	midgard.Port = *flag_port

	// Daemon mode: enter the realm of the daemons
	if *flag_daemon {
		helheim.MainDaemon()
		return
	}

	// Client mode: stay in the realm of the humans
	args := flag.Args()
	if len(args) > 1 {
		if args[0] == "add" || args[0] == "rm" || args[0] == "requeue" {
			ExpandFiles(args[1:])
		}
	}
	midgard.MainClient(args)
}

// Make file paths absolute.
func ExpandFiles(args []string) {
	wd, err := os.Getwd()
	helheim.Check(err)
	for i, arg := range args {
		if !path.IsAbs(arg) && FileExists(arg) {
			args[i] = path.Clean(wd + "/" + arg)
		}
	}
}

func FileExists(file string) bool {
	_, err := os.Stat(file)
	return err == nil
}
