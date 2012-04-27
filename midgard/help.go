package midgard

// This file implements the "help" command.

import (
	"fmt"
	"io"
)

// Store help for commands here
var help map[string]string = make(map[string]string)

func init() {
	help["help"] = `Display this help message`
	Api["help"] = Help
}

func Help(out io.Writer) error {
	fmt.Fprintln(out, `usage: `+Prog+` <command> [<args>]`)
	fmt.Fprint(out, `The available commands are:`)
	for name, _ := range Api {
		fmt.Fprint(out, "\n   "+fill(name)+" "+help[name])
	}
	return nil
}

// paste some spaces after the string for column alignment
func fill(str string) string {
	return str + "          "[:10-len(str)]
}
