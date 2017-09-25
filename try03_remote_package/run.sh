#!/usr/bin/env bash

REMOTE_PACKAGE='github.com/golang/example/hello'
go get $REMOTE_PACKAGE
exec "$GOPATH/bin/hello"
exec "hello"
