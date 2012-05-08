package midgard

// This file implements the client main function.
// Invoked whenever the user types the command. 
// The client merely forwards the CLI arguments
// to the daemon and returns the response to user.

import (
	"fmt"
	"net/rpc"
	"os"
	"os/exec"
	"os/user"
	"strings"
	"time"
)

// Args struct passed from client to daemon
type Args struct {
	CliArgs []string   // Command line args
	User    *user.User // User info
}

// Main loop for "client" mode (the normal mode).
// Simply passes the arguments to the daemon and
// displays the result.
func MainClient(args []string) {
	client := dialDaemon()
	var resp string
	user, err := user.Current()
	Check(err)
	err = client.Call("RPC.Call", &Args{args, user}, &resp)
	if err != nil {
		fmt.Fprint(os.Stderr, cleanup(err.Error()))
	}
	fmt.Print(cleanup(resp))
}

// cleanup newlines so string can be printed to stdout without redundant/missing newlines
func cleanup(str string) string {
	str = strings.Trim(str, "\n")
	if str != "" {
		return str + "\n"
	}
	return str
}

// Connect to the daemon for RPC communication.
// Starts the daemon if he's not yet running.
func dialDaemon() *rpc.Client {
	// try to call the daemon
	client, err := rpc.DialHTTP("tcp", "localhost"+Port)

	// if daemon does not seem to be running, start him.
	if SpawnDaemon {
		const SLEEP = 10e6 // nanoseconds
		if err != nil {
			forkDaemon()
			time.Sleep(SLEEP)
		}

		// try again to call the daemon,
		// give him some time to come up.
		trials := 0
		for err != nil && trials < 10 {
			client, err = rpc.DialHTTP("tcp", "localhost"+Port)
			time.Sleep(SLEEP)
			trials++
		}
	}

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	return client
}

// Start the daemon.
func forkDaemon() {
	executable, err := os.Readlink("/proc/self/exe")
	Check(err)
	cmd := exec.Command(executable, "-d")
	err = cmd.Start()
	Check(err)
}
