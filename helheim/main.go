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
	midgard.Help["add"] = "Add input file to job queue"

	midgard.Api["rm"] = Rm
	midgard.Help["rm"] = "Remove input file from queue"

	midgard.Api["status"] = Status
	midgard.Help["status"] = "Show queue status"

	midgard.Api["addnode"] = AddNodeAPI
	midgard.Help["addnode"] = "Add a compute node"

	midgard.Api["addgroup"] = AddGroupAPI
	midgard.Help["addgroup"] = "Add a group of users"

	midgard.Api["adduser"] = AddUserAPI
	midgard.Help["adduser"] = "Add a user to an existing group"

	midgard.Api["config"] = Setv
	midgard.Help["config"] = "Set configuration variables"

	midgard.Api["requeue"] = Requeue
	midgard.Help["requeue"] = "Kill a job and re-queue it for later"
}
