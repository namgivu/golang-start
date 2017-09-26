package main

/*
https://www.hackerrank.com/challenges/time-conversion
*/

import (
	"fmt"
	"github.com/namgivu/golang-start/hackerrank/util"
	"strconv"
)

const CHALLENGE_NAME = "ch170926_1346_am-pm-24h"

func main() {

	/*region io setting*/

	var REDIRECT_STDIO_2_FILE bool
	REDIRECT_STDIO_2_FILE = false
	REDIRECT_STDIO_2_FILE = true //turn this on to redirect stdin/stdout to file ref. TODO https://stackoverflow.com/q/46399395/248616

	/*region file path config*/
	var IO_HOME = util.StrFormat("{HACKERRANK_HOME}/{CHALLENGE_NAME}", "{HACKERRANK_HOME}", util.HACKERRANK_HOME, "{CHALLENGE_NAME}", CHALLENGE_NAME)
	var INPUT_FILE = util.StrFormat("{IO_HOME}/input.txt", "{IO_HOME}", IO_HOME)
	var OUTPUT_FILE = util.StrFormat("{IO_HOME}/output.txt", "{IO_HOME}", IO_HOME)
	/*endregion file path config*/

	/*region redirect stdio*/
	if REDIRECT_STDIO_2_FILE {
		inFile, outFile := util.RedirectStdio2File(INPUT_FILE, OUTPUT_FILE)
		defer inFile.Close()
		defer outFile.Close()
	}
	/*endregion redirect stdio*/

	/*endregion io setting*/

	var t string
	fmt.Scanf("%s", &t)

	hh := t[0:2]
	h,e := strconv.Atoi(hh)
	if e!=nil { panic(e) }

	ampm := t[len(t)-2:]
	if ampm == "PM" {
		h += 12
	}

	mm := t[3:5]
	ss := t[6:8]
	r := fmt.Sprintf("%02d:%02s:%02s", h, mm, ss)
	fmt.Println(r)
}
