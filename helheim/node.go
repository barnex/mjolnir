package helheim

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os/exec"
	"os/user"
)

// Compute node.
type Node struct {
	name    string
	err     error     // Error message configuring the node, if any.
	ssh     []string  // Command to login to this node. E.g.: {"ssh", "localhost"}.
	devices []*Device // GPUs in the node.
}

// Add a new compute node and return it as well.
func AddNode(name string, ssh ...string) error {
	node := &Node{name, nil, ssh, []*Device{}}
	nodes = append(nodes, node)
	node.Autoconf()
	return node.err
}

// API func to add new node.
func AddNodeAPI(out io.Writer, user *user.User, args []string) error {
	err := AddNode(args[0], args[1:]...)
	return err
}

func (n *Node) Autoconf() {
	// Ask for node auto config.
	var info NodeInfo
	bytes, err := n.Exec("", MUNINN) //TODO
	if err == nil {
		Check(json.Unmarshal(bytes, &info))
		Debug("muninn says: ", info)
	} else {
		info.ErrorString = fmt.Sprint(err, ": ", string(bytes))
	}

	// Store received config info.
	if info.ErrorString != "" {
		n.err = errors.New(info.ErrorString)
	}
	//Debug("len(info.Devices))", len(info.Devices))
	n.devices = make([]*Device, len(info.Devices))
	for i, dev := range info.Devices {
		//Debug("dev", i, dev)
		n.devices[i] = &Device{dev.Name, dev.TotalMem, false, false}
	}
}

// Execute a command on the node
func (n *Node) Exec(wd string, command string, args ...string) (output []byte, err error) {
	cmd := n.Cmd(wd, command, args...)
	output, err = cmd.CombinedOutput()
	Debug(string(output))
	return
}

// Prepare a command for execution on the node.
// Prefixes the command line with the appropriate ssh stanza for this node.
func (n *Node) Cmd(wd string, command string, args ...string) *exec.Cmd {
	allArgs := append(append(n.ssh[1:], command), args...)
	cmd := exec.Command(n.ssh[0], allArgs...)
	cmd.Dir = wd
	Debug("exec: ", n.ssh[0], allArgs)
	return cmd
}

// API func, prints node info.
func Nodes(out io.Writer) error {
	for _, n := range nodes {
		fmt.Fprintln(out, n)
		for _, d := range n.devices {
			fmt.Fprintln(out, "\t", d)
		}
	}
	return nil
}

func (n *Node) String() string {
	if n.err == nil {
		return n.name
	}
	return n.name + ": " + n.err.Error()
}
