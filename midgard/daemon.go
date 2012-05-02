package midgard

// This file implements the Remote Procedure Call between
// the daemon and client front-end
// The client forks a daemon 
// if none is yet running and sends RPC calls to it.

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"net"
	"net/http"
	"net/rpc"
	"os/user"
	"reflect"
)

// Start serving RPC calls from client instances.
func Listen() {
	rpc.Register(&RPC{})
	rpc.HandleHTTP()
	conn, err := net.Listen("tcp", Port)
	if err != nil {
		Err("listen error:", err)
	}
	Debug("Listening on port " + Port)
	http.Serve(conn, nil)
	//TODO: log errors.
}

// Dummy type to define RPC methods on
type RPC struct {
}

// RPC-exported function used for normal operation mode.
// The command-line arguments are passed (e.g. "play jazz")
// and a response to the user is returned in *resp.
// Here, run-time reflection is used to match the user command
// to a method on the API type.
func (rpc RPC) Call(argz Args, resp *string) (err error) {

	args := argz.CliArgs
	usr := argz.User

	Debug("midgard: aquire lock")
	Lock.Lock()
	defer Lock.Unlock()
	defer Debug("midgard: release lock")

	Debug("ServerRPC.Call", usr, args)

	if len(args) == 0 {
		args = []string{"help"}
	}

	cmd := args[0]  // first arg is command
	args = args[1:] // rest are arguments

	// lookup function in API map
	f := Api[cmd]
	if f == nil {
		err = errors.New(ProgName + ": '" + cmd + "' is not a " + ProgName + " command. See " + ProgName + " help.")
		return
	}

	// call function
	var out_ bytes.Buffer
	out := &out_
	switch fnc := f.(type) {
	default:
		panic(errors.New(fmt.Sprint("midgard: unsupported func type for ", cmd, " : ", reflect.TypeOf(f))))
	case func(string):
		if len(args) != 0 {
			err = NewError(cmd, ": needs one arugment")
		} else {
			fnc(args[0])
		}
	case func(io.Writer) error:
		if len(args) > 0 {
			err = NewError(cmd, ": does not take arugments")
		} else {
			err = fnc(out)
		}
	case func(io.Writer, string) error:
		if len(args) != 1 {
			err = NewError(cmd, ": needs one argument")
		} else {
			err = fnc(out, args[0])
		}
	case func(io.Writer, []string) error:
		if len(args) == 0 {
			err = NewError(cmd, ": needs argument")
		} else {
			err = fnc(out, args)
		}
	case func(io.Writer, *user.User, []string) error:
		if len(args) == 0 {
			err = NewError(cmd, ": needs argument")
		} else {
			err = fnc(out, usr, args)
		}
	}
	*resp = string(out.Bytes())
	return
}
