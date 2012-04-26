package midgard

// This file implements the "help" command.

import ()

// Store help for commands here
var help map[string]string = make(map[string]string)

func init() {
	help["help"] = `Display this help message`
	Api["help"] = (*Server).Help
}

func Help() (resp, err string) {
	resp = `usage: ` + Prog + ` <command> [<args>]

The available commands are:`
	for name, _ := range Api {
		resp += "\n   " + fill(name) + " " + help[name]
	}
	return
}

// paste some spaces after the string for column alignment
func fill(str string) string {
	return str + "          "[:10-len(str)]
}
