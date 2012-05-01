package helheim

import (
	"encoding/json"
	"fmt"
	"io"
	"os/exec"
)

// Compute node.
type Node struct {
	errorString string    // Error message configuring the node, if any.
	ssh         []string  // Command to login to this node. E.g.: {"ssh", "localhost"}.
	devices     []*Device // GPUs in the node.
}

// Compute device.
type Device struct {
	name     string
	totalMem int64
	busy     bool
	drain    bool
}

func (d *Device) String() string {
	return fmt.Sprint(d.name, " ", d.Megabytes(), "MB")
}

// Total memory in megabytes.
func (d *Device) Megabytes() int {
	return int(d.totalMem / (1024 * 1024))
}

func AddNode(ssh []string) {
	node := &Node{"", ssh, []*Device{}}
	nodes = append(nodes, node)
	node.Autoconf()
}

func (n *Node) Autoconf() {
	// Ask for node auto config.
	var info NodeInfo
	bytes, err := n.Exec("/home/arne/go/bin/muninn") //TODO
	if err == nil {
		Check(json.Unmarshal(bytes, &info))
		Debug("muninn says: ", info)
	} else {
		info.ErrorString = fmt.Sprint(err, ": ", string(bytes))
	}

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
	allArgs := append(append(n.ssh[1:], command), args...)
	cmd := exec.Command(n.ssh[0], allArgs...)
	Debug("exec: ", n.ssh[0], allArgs)
	output, err = cmd.CombinedOutput()
	Debug(string(output))
	return
}

// API func, prints node info.
func Nodes(out io.Writer) error {
	for _, n := range nodes {
		fmt.Fprintln(out, n.ssh, n.errorString)
		for _, d := range n.devices {
			fmt.Fprintln(out, "\t", d)
		}
	}
	return nil
}
