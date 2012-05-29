package helheim

import (
	"errors"
	"strconv"
	"strings"
)

func ParseStringFlag(args []string, flag string) (otherargs []string, value string, err error) {

	valuei := -1
	for i, arg := range args {
		if strings.HasPrefix(arg, flag) {
			if i == len(args)-1 {
				err = errors.New(flag + "needs argument")
				return
			}
			value = args[i+1]
			valuei = i
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

func ParseIntFlag(args []string, flag string) (otherargs []string, value int, err error) {
	str := ""
	otherargs, str, err = ParseStringFlag(args, flag)
	if err != nil {
		return
	}
	if str != ""{
		value, err = strconv.Atoi(str)
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
