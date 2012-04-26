package midgard

// This file implements the Remote Procedure Call between
// the daemon and client front-end

import (
	"errors"
	"fmt"
	"reflect"
)

// Aliased type to define RPC methods on.
type RPC struct {
	server *Server
}

// RPC-exported function used for normal operation mode.
// The command-line arguments are passed (e.g. "play jazz")
// and a response to the user is returned in *resp.
// Here, run-time reflection is used to match the user command
// to a method on the API type.
func (rpc RPC) Call(args []string, resp *string) (err error) {
	Debug("ServerRPC.Call", args)

	if len(args) == 0 {
		args = []string{"help"}
	}

	cmd := args[0]  // first arg is command (e.g.: "play")
	args = args[1:] // rest are arguments (e.g.: "jazz")

	// lookup function in API map
	f := Api[cmd]
	if f == nil {
		err = errors.New(Prog + ": '" + cmd + "' is not a " + Prog + " command. See " + Prog + " help.")
		return
	}

	// call function
	errstr := ""
	switch fnc := f.(type) {
	default:
		panic(reflect.TypeOf(f))
	case func() (string, string):
		if len(args) > 0 {
			errstr = fmt.Sprint(cmd, ": does not take arugments")
		} else {
			*resp, errstr = fnc()
		}
	case func(string) (string, string):
		if len(args) != 1 {
			errstr = fmt.Sprint(cmd, ": needs one argument")
		} else {
			*resp, errstr = fnc(args[0])
		}
	case func([]string) (string, string):
		if len(args) == 0 {
			errstr = fmt.Sprint(cmd, ": needs argument")
		} else {
			*resp, errstr = fnc(args)
		}
	}
	if errstr != "" {
		err = errors.New(errstr)
	}
	return
}

