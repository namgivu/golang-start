#!/usr/bin/env bash

#Objective: we build go code to go executable file at $buildOutput

#get SCRIPT_HOME=executed script's path, containing folder, cd & pwd to get container path
s=$BASH_SOURCE ; s=$(dirname "$s") ; s=$(cd "$s" && pwd) ; export SCRIPT_HOME=$s

#region do build
  NAME=$(basename "$SCRIPT_HOME")
  cd $SCRIPT_HOME
  go build -o $NAME main.go
  cd -- #go back where we were
#endregion do build

buildOutput="$SCRIPT_HOME/$NAME"
echo "
Built output at
$buildOutput
"
