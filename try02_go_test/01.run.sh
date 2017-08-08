#!/usr/bin/env bash

#get SCRIPT_HOME=executed script's path, containing folder, cd & pwd to get container path
s=$BASH_SOURCE ; s=$(dirname "$s") ; s=$(cd "$s" && pwd) ; export SCRIPT_HOME=$s

GOPATH_SRC="$GOPATH/src/"
PACKAGE_PATH=${SCRIPT_HOME/#$GOPATH_SRC/}
go test $PACKAGE_PATH
