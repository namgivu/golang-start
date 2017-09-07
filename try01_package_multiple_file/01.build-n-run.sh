#!/usr/bin/env bash

#get SCRIPT_HOME=executed script's path, containing folder, cd & pwd to get container path
s=$BASH_SOURCE ; s=$(dirname "$s") ; s=$(cd "$s" && pwd) ; export SCRIPT_HOME=$s

#build output
NAME='try01'

#build it
cd $SCRIPT_HOME
go build -o $NAME
cd --

#run it
exec "$SCRIPT_HOME/$NAME"
