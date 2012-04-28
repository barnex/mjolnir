package helheim

import (
	"fmt"
	"io"
)

// Compute node.
type Node struct {
	hostname string
	device   []*Device
}

// Compute device.
type Device struct {
	id    int
	busy  bool
	drain bool
}

func AddNode(hostname string) {
	node := &Node{hostname, []*Device{}}
	nodes = append(nodes, node)
}

// API func, prints node info.
func Nodes(out io.Writer) error {
	for _, n := range nodes {
		fmt.Fprint(out, n)
	}
	return nil
}
