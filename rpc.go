package main

// This file implements the Remote Procedure Call between
// the daemon and client front-end

import (
	"errors"
	"fmt"
	"net"
	"net/http"
	"net/rpc"
	"reflect"
)

// Start serving RPC calls from client instances.
func (player *Server) serveRPC() {
	rpc.Register(player.RPC())
	rpc.HandleHTTP()
	conn, err := net.Listen("tcp", port)
	if err != nil {
		Err("listen error:", err)
	}
	Debug("Listening on port " + port)
	http.Serve(conn, nil)
	//TODO: log errors.
}

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
	f := api[cmd]
	if f == nil {
		err = errors.New(PROG + ": '" + cmd + "' is not a " + PROG + " command. See " + PROG + " help.")
		return
	}

	// call function
	errstr := ""
	switch fnc := f.(type) {
	default:
		panic(reflect.TypeOf(f))
	case func(*Server) (string, string):
		if len(args) > 0 {
			errstr = fmt.Sprint(cmd, ": does not take arugments")
		} else {
			*resp, errstr = fnc(rpc.server)
		}
	case func(*Server, string) (string, string):
		if len(args) != 1 {
			errstr = fmt.Sprint(cmd, ": needs one argument")
		} else {
			*resp, errstr = fnc(rpc.server, args[0])
		}
	case func(*Server, []string) (string, string):
		if len(args) == 0 {
			errstr = fmt.Sprint(cmd, ": needs argument")
		} else {
			*resp, errstr = fnc(rpc.server, args)
		}
	}
	if errstr != "" {
		err = errors.New(errstr)
	}
	return
}

var api map[string]interface{} = make(map[string]interface{})
