package helheim

// Compute device.
type Device struct {
	host  *Node
	id    int
	busy  bool
	drain bool
}
