#!/usr/bin/env bash

#get SCRIPT_HOME=executed script's path, containing folder, cd & pwd to get container path
s=$BASH_SOURCE ; s=$(dirname "$s") ; s=$(cd "$s" && pwd) ; export SCRIPT_HOME=$s

#region do install
curDir=`pwd`

  NAME='try00b'
  cd $SCRIPT_HOME
  go install -o $NAME main.go

cd $curDir
#endregion do install

buildOutput="$GOPATH/bin/$NAME"
echo "
Built output at
$buildOutput #TODO We do NOT see the output here, why?
"
