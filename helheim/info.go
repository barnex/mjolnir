package helheim

// NodeInfo struct is JSON-encoded by muninn
// and sent for node auto-config.
type NodeInfo struct {
	CudaError   int          // Passes a CUDA error, if any
	ErrorString string       // Passes message of any error
	Devices     []DeviceInfo // Passes device info
}

type DeviceInfo struct {
	Name     string
	TotalMem int64
}
