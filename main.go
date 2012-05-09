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
	args := flag.Args()
	if len(args) > 1 {
		if args[0] == "add" || args[0] == "rm"{
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
		if !path.IsAbs(arg) && FileExists(arg){
			args[i] = path.Clean(wd + "/" + arg)
		}
	}
}

func FileExists(file string)bool{
		_, err := os.Stat(file)
		return err == nil
}
