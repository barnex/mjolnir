# Mjolnir

Mjolnir is a queue manager for GPU clusters, targeted at mumax.
It has been superseeded by mumax3-server.

## Cofiguration:

On the head node, run:

	mjolnir -d

which starts the mjolnir daemon. This listens for non-daemon mjolnir
processes over TCP.

Then add compute nodes and specify a login command:

	mjolnir addnode mybox ssh 127.0.0.1
	mjolnir addnode bigserver ssh 192.168.0.1


	mjolnir nodes

Prints node and GPU info. E.g.:

		mybox
	 		NVS 3100M 511MB free
		bigserver
			GTX 680 2047MB busy
			GTX 680 2047MB free

The node's GPUs are automatically discovered. This requires the
executable "muninn" (goinstallable as mjolnir/muninn) to be in the
compute nodes' $PATH. Alternatively the path to muninn can be set with:

	mjolnir config muninn $GOPATH/bin/muninn

Then add compute groups and grant them a number of GPUs:

	mjolnir addgroup group1 8 # group1 gets 8 GPUs
	mjolnir addgroup group2 4 # group2 gets 4 GPUs

Then add users to the groups and also grant them a number of GPUs:

	mjolnir adduser jack group1 2 # jack gets 2 of group1's 8 GPUs
	mjolnir adduser jill group2 4


	mjolnir groups
	mjolnir users

will show group and user info.

The specified total number of GPUs does not need to correspond to the
total number of GPUs actually present in the group or cluster. They are
only used in a relative way. GPUs are first allocated per group, using
relative shares, and then per user inside the group.

## Use:

	mjolnir help

Prints all available commands.

Adding input files to the queue is done with the "add" command.
Wildcards can be used as they are expanded by the shell. Priorities can
be set with the -pr flag. Files with higher priority will run first. If
the same file is added twice, it will only be in the queue once.
However, adding an already present file may be used to change its
priority.

	mjolnir add file1.py
	mjolnir add *.py
	mjolnir add file1.py -pr 10 # changes priority of file1.py
	mjolnir add file1.py -gpus 4                  # run on 4 GPUs
	mjolnir add file1.py -exec /bin/my_executable # use custom executable
	mjolnir add file1.py -wall 1h                 # set maximum walltime. default 24h.

Files can be removed from the queue with

	mjolnir rm file1.py

Removing a running file will kill the compute job.

Watch the queue status with:

	mjolnir status
