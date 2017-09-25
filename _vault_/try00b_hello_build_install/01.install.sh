#!/usr/bin/env bash

##Please put the project folder under $GOPATH/src/github.com/$YOUR_GITHUB_USER/

#get SCRIPT_HOME=executed script's path, containing folder, cd & pwd to get container path
s=$BASH_SOURCE ; s=$(dirname "$s") ; s=$(cd "$s" && pwd) ; export SCRIPT_HOME=$s

#region do install
NAME=$(basename "$SCRIPT_HOME") #TODO How to name the installed output?

cd $SCRIPT_HOME

GOPATH_SRC="$GOPATH/src/"
PACKAGE_PATH="${SCRIPT_HOME/#$GOPATH_SRC/}" #remove absolute part in the path i.e. $GOPATH/src ref. https://stackoverflow.com/a/24347937/248616
sh="go install $PACKAGE_PATH" #go install syntax ref. https://golang.org/doc/code.html#Command

echo "
GOPATH_SRC=$GOPATH_SRC
PACKAGE_PATH=$PACKAGE_PATH

Install command:
$sh
"

eval $sh
#endregion do install

#output
buildOutput="$GOPATH/bin/$NAME"
echo "
Built output at
$buildOutput

You can run above full path, or just call
$NAME
"
