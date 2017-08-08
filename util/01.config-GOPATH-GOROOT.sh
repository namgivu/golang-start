#!/usr/bin/env bash

#what are they        ref. https://stackoverflow.com/a/10847122/248616
#official guide       ref. https://golang.org/doc/code.html
#more friendly guide  ref. https://github.com/alco/gostart#motivation

GOPATH="$HOME/gows" #gows ie. go workspace
GOROOT=`which go` #golang installed in ubuntu, just use that path

BASHRC="$HOME/.bashrc"

#set up GOPATH
mkdir -p $GOPATH

#register path via .bashrc
echo $BASHRC
echo
cat $BASHRC

cat >> $BASHRC <<EOL

#config GOROOT for golang
export GOROOT="$GOROOT"

#config GOPATH for golang
export GOPATH="$GOPATH"
export PATH="\$GOPATH/bin:\$PATH"
EOL

echo '---'
cat $BASHRC

#reload .bashrc
source $BASHRC
echo "
GOROOT=$GOROOT
GOPATH=$GOPATH
"
