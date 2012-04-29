package main

import (
	cu "cuda/driver"
	"encoding/json"
	"fmt"
	. "mjolnir/helheim"
	"os"
)

func main() {
	/*defer func(){
		err := recover()
		if err != nil{

		}
	}()
	*/
	var info NodeInfo
	cu.Init(0)
	NDev := cu.DeviceGetCount()
	info.Devices = make([]DeviceInfo, NDev)
	for i := range info.Devices {
		dev := cu.DeviceGet(i)
		info.Devices[i].Name = dev.Name()
	}
	bytes, err := json.Marshal(info)
	Check(err)
	_, err2 := os.Stdout.Write(bytes)
	fmt.Println()
	Check(err2)
}
