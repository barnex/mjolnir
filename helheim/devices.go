package helheim

// Compute device.
type Device struct {
	host  *Node
	id    int
	busy  bool
	drain bool
}

// Compute node.
type Node struct {
	hostname string
	device   []*Device
}
