#!/usr/bin/env bash

#latest version ref. https://golang.org/dl/#stable
#install guide ref. https://stackoverflow.com/a/43512454/248616


#install
sudo add-apt-repository ppa:longsleep/golang-backports
sudo apt update
sudo apt install -y golang-go

#clean up the unused
sudo apt -y autoremove

#aftermath check
echo
which go
go version

#set up GOPATH and GOROOT
  #get SCRIPT_HOME=executed script's path, containing folder, cd & pwd to get container path
  s=$BASH_SOURCE ; s=$(dirname "$s") ; s=$(cd "$s" && pwd) ; export SCRIPT_HOME=$s

  source "$SCRIPT_HOME/01.config-GOPATH-GOROOT.sh"
