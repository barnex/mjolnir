package helheim

import (
	"fmt"
	"io"
	"mjolnir/midgard"
	"runtime"
)

func init() {
	midgard.Api["version"] = Version
	midgard.Help["version"] = "Print version info"
}

func Version(out io.Writer) (err error) {
	fmt.Fprintln(out, `Mjǫlnir version 0.0.1`)
	fmt.Fprintln(out, "Go version", runtime.Version())
	return
}
