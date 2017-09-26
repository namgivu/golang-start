package main

/*
https://www.hackerrank.com/challenges/staircase
*/

import (
	"fmt"
	"github.com/namgivu/golang-start/hackerrank/util"
	"strings"
)

const CHALLENGE_NAME = "ch170926_1200_staircase"

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

	var N int
	fmt.Scanf("%d", &N)
	r := ""
	for i:=1; i<=N; i++ {
		spaces := N-i
		stars  := i
		s := ""
		for j:=1; j<=spaces; j++ {
			s += " "
		}
		for j:=1; j<=stars; j++ {
			s += "#"
		}
		r += s+"\n"
	}
	r = strings.TrimRight(r, "\n")
	fmt.Print(r)
}
