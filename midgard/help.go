package midgard

// This file implements the "help" command.

import (
	"fmt"
	"io"
)

func init() {
	Help["help"] = `Display this help message`
	Api["help"] = PrintHelp
}

func PrintHelp(out io.Writer) error {
	fmt.Fprintln(out, `usage: `, ProgName, ` <command> [<args>]`)
	fmt.Fprint(out, `The available commands are:`)
	for name, _ := range Api {
		fmt.Fprint(out, "\n   ", fill(name), " ", Help[name])
	}
	return nil
}

// paste some spaces after the string for column alignment
func fill(str string) string {
	return str + "          "[:10-len(str)]
}
