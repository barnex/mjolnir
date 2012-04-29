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

	defer func() {
		err := recover()
		if err != nil {
			info.ErrorString = fmt.Sprint(err)
			if cudaErr, ok := err.(cu.Result); ok {
				info.CudaError = int(cudaErr)
			}
		}

		bytes, err2 := json.Marshal(info)
		Check(err2)
		_, err3 := os.Stdout.Write(bytes)
		fmt.Println()
		Check(err3)
	}()

	cu.Init(7)
	NDev := cu.DeviceGetCount()
	info.Devices = make([]DeviceInfo, NDev)
	for i := range info.Devices {
		dev := cu.DeviceGet(i)
		info.Devices[i].Name = dev.Name()
	}
}
