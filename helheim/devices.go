package helheim

import (
	"fmt"
	"io"
	"os/exec"
)

// Compute node.
type Node struct {
	ssh []string
	device   []*Device
}

// Compute device.
type Device struct {
	id    int
	busy  bool
	drain bool
}

func AddNode(ssh []string) {
	node := &Node{ssh, []*Device{}}
	nodes = append(nodes, node)
	node.Autoconf()
}

func(n*Node)Autoconf(){
	n.Exec("hostname")
}

// Execute a command on the node
func(n*Node)Exec(command string, args ...string) (output[]byte, err error){
	Debug("exec", command, args)
	Debug(args)
	Debug(n.ssh[0], append(append(n.ssh[1:], command), args...))	
	cmd := exec.Command(n.ssh[0], append(n.ssh[1:], args...)...)	
	Debug(cmd)
	output, err = cmd.CombinedOutput()
	Check(err)
	Debug(string(output))
	return
}

// API func, prints node info.
func Nodes(out io.Writer) error {
	for _, n := range nodes {
		fmt.Fprint(out, n)
	}
	return nil
}
