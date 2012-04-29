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
	cu.Init(0)
	NDev := cu.DeviceGetCount()
	info := make([]DeviceInfo, NDev)
	for i := range info {
		dev := cu.DeviceGet(i)
		info[i].Name = dev.Name()
	}
	bytes, err := json.Marshal(info)
	Check(err)
	_, err2 := os.Stdout.Write(bytes)
	fmt.Println()
	Check(err2)
}
