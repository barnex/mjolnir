package helheim

import (
	"errors"
	"io"
	"os/user"
)

// executable
var (
	executable []string = []string{"mumax2", "-s"} // Executable and arguments to run input files
	muninn              = "muninn"                 // To be replaced by full path to muninn.
)

func Configure() {
}

func Setv(out io.Writer, usr *user.User, args []string) error {
	key := args[0]
	val := args[1:]

	switch key {
	default:
		return errors.New("invalid key: " + key)
	case "mumax2":
		executable = val
	case "muninn":
		muninn = val[0]
	}
	return nil
}
