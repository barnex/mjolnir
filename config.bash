#! /bin/bash
killall mjolnir
mjolnir -d & sleep 1s;
mjolnir config muninn $GOPATH/bin/muninn
mjolnir addnode localhost ssh localhost
mjolnir nodes
