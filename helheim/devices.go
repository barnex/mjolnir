package helheim

import (
	"encoding/json"
	"fmt"
	"io"
	"os/exec"
)

// Compute node.
type Node struct {
	errorString string // Error message configuring the node, if any.
	ssh     []string // Command to login to this node. E.g.: {"ssh", "localhost"}.
	devices []*Device // GPUs in the node.
}

// Compute device.
type Device struct {
	name string
	totalMem int64
	busy  bool
	drain bool
}

func(d*Device)String()string{
	return fmt.Sprint(d.name, " ", d.Megabytes(), "MB")
}

func(d*Device)Megabytes()int{
	return int(d.totalMem / (1024*1024))
}

func AddNode(ssh []string) {
	node := &Node{"", ssh, []*Device{}}
	nodes = append(nodes, node)
	node.Autoconf()
}

func (n *Node) Autoconf() {
	// Ask for node auto config.
	bytes, err := n.Exec("/home/arne/go/bin/muninn") //TODO
	Check(err)
	var info NodeInfo
	Check(json.Unmarshal(bytes, &info))
	Debug("muninn says: ", info)
	
	// Store received config info.
	n.errorString = info.ErrorString
	Debug("len(info.Devices))", len(info.Devices))
	n.devices = make([]*Device, len(info.Devices))
		for i, dev := range info.Devices {
			Debug("dev", i, dev)
			n.devices[i] = &Device{dev.Name, dev.TotalMem, false, false}
		}
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
		fmt.Fprintln(out, n.ssh)
			for _, d := range n.devices{
				fmt.Fprintln(out, "\t", d)
			}
	}
	return nil
}
