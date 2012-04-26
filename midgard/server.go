package midgard

// This file implements the queue state.
// The player communicates over RPC (Remote Procedure Call)
// with the client. The client forks a daemon 
// if none is yet running and sends RPC calls to it.

import (
//"sync"
)

type Server struct {
	port string // default RPC port
}

// Constructor
func NewServer() *Server {
	p := new(Server)
	p.init()
	return p
}

// Wraps the player in an RPC to expose methods available to the RPC server.
func (p *Server) RPC() RPC {
	return RPC{p}
}

func (p *Server) init() {
	Debug("server initialized")
}

// Main loop for daemon mode
func (p *Server) Daemon() {
	// TODO: heartbeat here
	p.serveRPC()
}
