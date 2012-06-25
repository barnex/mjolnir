package helheim

import (
	"fmt"
	"os"
	"strings"
	"time"
)

var logcount int

func checkLog() {
	logcount++
	// ever x heartbeats
	if logcount == 12 {
		logcount = 0
	} else {
		return
	}

	logTemperature()
}

var tempFile *os.File

func logTemperature() {

	var err error
	if tempFile == nil {
		tempFile, err = os.OpenFile("temperature.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0777)
		fmt.Fprintln(tempFile, "# unix time (s)	device temp (C)")
		fmt.Fprint(tempFile, "#")
		for _, n := range nodes {
			for d := range n.devices {
				fmt.Fprint(tempFile, n.name,":",d,"\t")
			}
		}
		fmt.Fprintln(tempFile)
		fmt.Fprintln(tempFile) // newline interrupts gnuplot trace.
	}
	Check(err)
	now := time.Now()
	fmt.Fprint(tempFile, now.Unix())

	for _, n := range nodes {
		for d := range n.devices {

			bytes, err := n.Exec("", "nvidia-smi", "-q", fmt.Sprint("--id=", d))

			if err != nil {
				Debug(err, ": ", string(bytes))
				return
			}
			resp := string(bytes)
			lines := strings.Split(resp, "\n")
			for _, l := range lines {
				l := strings.Trim(l, " ")
				if strings.HasPrefix(l, "Gpu") && strings.HasSuffix(l, "C") {
					l = l[26 : len(l)-2]
					fmt.Fprint(tempFile, "\t", l)
				}
			}
		}
	}
	fmt.Fprintln(tempFile, "\t# ", now.Format(time.ANSIC))
}
