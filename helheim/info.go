package helheim

type NodeInfo struct {
	CudaError   int
	ErrorString string
	Devices     []DeviceInfo
}

type DeviceInfo struct {
	Name     string
	TotalMem int64
}
