#!/usr/bin/env bash

#get SCRIPT_HOME=executed script's path, containing folder, cd & pwd to get container path
s=$BASH_SOURCE ; s=$(dirname "$s") ; s=$(cd "$s" && pwd) ; export SCRIPT_HOME=$s

#region do build
curDir=`pwd`

  NAME='try00b'
  cd $SCRIPT_HOME
  go build -o $NAME main.go

cd $curDir
#endregion do build

buildOutput="$SCRIPT_HOME/$NAME"
echo "
Built output at
$buildOutput
"
