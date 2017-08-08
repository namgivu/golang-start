#!/usr/bin/env bash

#get SCRIPT_HOME=executed script's path, containing folder, cd & pwd to get container path
s=$BASH_SOURCE ; s=$(dirname "$s") ; s=$(cd "$s" && pwd) ; export SCRIPT_HOME=$s

#region do install
NAME=$(basename "$SCRIPT_HOME") #TODO How to build with this name as the output?

cd $SCRIPT_HOME

GOPATH_SRC="$GOPATH/src/"
PACKAGE_PATH="${SCRIPT_HOME/#$GOPATH_SRC/}" #replace substring ref. https://stackoverflow.com/a/24347937/248616

go install $PACKAGE_PATH #go install syntax ref. https://golang.org/doc/code.html#Command
#endregion do install

#output
buildOutput="$GOPATH/bin/$NAME"
echo "
Built output at
$buildOutput

You can run above full path, or just call
$NAME
"
