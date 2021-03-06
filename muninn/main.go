package main

import (
	cu "cuda/driver"
	"encoding/json"
	"fmt"
	. "mjolnir/helheim"
	"os"
)

func main() {
	var info NodeInfo

	// If an error occurs, send it in the NodeInfo
	defer func() {
		err := recover()
		if err != nil {
			info.ErrorString = fmt.Sprint(err)
			if cudaErr, ok := err.(cu.Result); ok {
				info.CudaError = int(cudaErr)
			}
		}

		// Send the info no matter what
		bytes, err2 := json.Marshal(info)
		Check(err2)
		_, err3 := os.Stdout.Write(bytes)
		fmt.Println()
		Check(err3)
	}()

	cu.Init(0)
	NDev := cu.DeviceGetCount()
	info.Devices = make([]DeviceInfo, NDev)
	for i := range info.Devices {
		dev := cu.DeviceGet(i)
		info.Devices[i].Name = dev.Name()
		info.Devices[i].TotalMem = dev.TotalMem()
	}
}
