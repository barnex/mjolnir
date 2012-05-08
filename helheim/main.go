package helheim

import (
	"mjolnir/midgard"
)

func MainDaemon() {

	Configure()
	initMidgard()

	FillNodes()

	go RunHeartbeat()

	// Start listening for commands
	midgard.Listen()
}

func initMidgard() {
	midgard.Lock = &lock

	midgard.Api["version"] = Version
	midgard.Help["version"] = "Print version info"

	midgard.Api["groups"] = Groups
	midgard.Help["groups"] = "Print group info"

	midgard.Api["next"] = PrintNext
	midgard.Help["next"] = "Print next job to run"

	midgard.Api["users"] = Users
	midgard.Help["users"] = "Print user info"

	midgard.Api["nodes"] = Nodes
	midgard.Help["nodes"] = "Print node info"

	midgard.Api["add"] = Add
	midgard.Help["add"] = "Add job"

	midgard.Api["status"] = Status
	midgard.Help["status"] = "Show queue status"

	midgard.Api["addnode"] = AddNodeAPI
	midgard.Help["addnode"] = "Add a compute node"

	midgard.Api["addgroup"] = AddGroupAPI
	midgard.Help["addgroup"] = "Add a user group"
}
