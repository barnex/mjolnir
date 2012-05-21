package helheim

import (
	"errors"
	"strconv"
	"strings"
)

func ParseFlag(args []string, flag string) (otherargs []string, value int, err error) {

	valuei := -1
	for i, arg := range args {
		if strings.HasPrefix(arg, flag) {
			if i == len(args)-1 {
				err = errors.New(flag + "needs argument")
				return
			}
			value, err = strconv.Atoi(args[i+1])
			valuei = i
			if err != nil {
				return
			}
			break
		}
	}
	// remove flag from args list
	if valuei != -1 {
		otherargs = append(args[:valuei], args[valuei+2:]...)
	} else {
		otherargs = args
	}
	return
}

func CheckNoMoreFlags(args []string) error {
	for _, arg := range args {
		if strings.HasPrefix(arg, "-") {
			return errors.New("invalid flag: " + arg)
		}
	}
	return nil
}
