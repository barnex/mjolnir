package helheim

import (
	"encoding/json"
	"fmt"
	"io"
	"os/exec"
)

// Compute node.
type Node struct {
	ssh     []string
	devices []*Device
}

// Compute device.
type Device struct {
	busy  bool
	drain bool
}

func AddNode(ssh []string) {
	node := &Node{ssh, []*Device{}}
	nodes = append(nodes, node)
	node.Autoconf()
}

func (n *Node) Autoconf() {
	bytes, err := n.Exec("/home/arne/go/bin/muninn") //TODO
	Check(err)
	var info NodeInfo
	err = json.Unmarshal(bytes, &info)
	Check(err)
	Debug("muninn says: ", info)
	//	n.devices = make([]*Device, len(info))
	//	for i := range n.devices {
	//		n.devices[i] = &Device{info[i], false, false}
	//	}
}

// Execute a command on the node
func (n *Node) Exec(command string, args ...string) (output []byte, err error) {
	cmd := exec.Command(n.ssh[0], append(append(n.ssh[1:], command), args...)...)
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
