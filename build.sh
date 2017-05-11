#!/bin/bash

export PACKAGE=canary
export DEST_PATH=/opt/$PACKAGE


env GOOS=linux GOARCH=amd64 go build -v $PACKAGE.go
scp $PACKAGE root@107.170.25.71:$DEST_PATH

