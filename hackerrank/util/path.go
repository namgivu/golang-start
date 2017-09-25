package util

import (
	"os"
)

var GOPATH							= os.Getenv("GOPATH")
var GITHUB_PROJECT_PATH	= "github.com/namgivu/golang-start"
var HACKERRANK_HOME			= StrFormat("{GOPATH}/src/{GITHUB_PROJECT_PATH}/hackerrank", "{GOPATH}", GOPATH, "{GITHUB_PROJECT_PATH}", GITHUB_PROJECT_PATH)
