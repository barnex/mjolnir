package helheim

import (
	"fmt"
	"io"
	"runtime"
)

func Version(out io.Writer) (err error) {
	fmt.Fprintln(out, `Mjǫlnir version 0.3`)
	fmt.Fprintln(out, "Go version", runtime.Version())
	return
}
