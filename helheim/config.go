package helheim

import (
	"errors"
	"io"
	"os/user"
)

// executable
var (
	executable []string = []string{"mumax2", "-s"}              // Executable and arguments to run input files
	muninn              = "muninn"                              // To be replaced by full path to muninn.
	translate           = []string{"/diskless/home/", "/home/"} // Translate input paths like this
)

func Configure() {
}

// Set a config key-value pair
func Setv(out io.Writer, usr *user.User, args []string) error {
	key := args[0]
	val := args[1:]

	switch key {
	default:
		return errors.New("invalid key: " + key)
	case "executable":
		executable = val
	case "muninn":
		muninn = val[0]
	case "translate":
		translate = []string{val[0], val[1]}
	}
	return nil
}
