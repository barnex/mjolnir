package main

import (
	cu "cuda/driver"
)

type NodeInfo struct {
	HostName string
	NDevice  int
	Devices  []Device
}

type Device struct {
	Name     string
	TotalMem int64
}

func main() {
	var info NodeInfo
	info.HostName = os.Hostname()
	NDev := cu.DeviceGetCount()

}
