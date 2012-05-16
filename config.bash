#! /bin/bash
killall mjolnir
mjolnir -d & sleep 1s;
mjolnir config muninn $GOPATH/bin/muninn
mjolnir config executable /home/arne/mumax2.git/bin/mumax2 -s
mjolnir addnode localhost ssh localhost
mjolnir addgroup local 1
mjolnir adduser arne local 1
mjolnir nodes
