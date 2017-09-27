package main

/*
https://www.hackerrank.com/challenges/compare-the-triplets
*/

import (
	"fmt"
	"github.com/namgivu/golang-start/hackerrank/util"
)

const CHALLENGE_NAME	= "ch170925_1733_triplet-compare"
const ARCHIVE_PATH		= "archive"

func main() {

	/*region io setting*/

	var REDIRECT_STDIO_2_FILE bool
	REDIRECT_STDIO_2_FILE = false
	REDIRECT_STDIO_2_FILE = true //turn this on to redirect stdin/stdout to file ref. https://stackoverflow.com/q/46399395/248616

	/*region file path config*/
	var IO_HOME = util.StrFormat("{HACKERRANK_HOME}/{ARCHIVE_PATH}/{CHALLENGE_NAME}", "{HACKERRANK_HOME}", util.HACKERRANK_HOME, "{ARCHIVE_PATH}", ARCHIVE_PATH, "{CHALLENGE_NAME}", CHALLENGE_NAME)
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

	a := make([]int, 3)
	fmt.Scanf("%d %d %d", &a[0], &a[1], &a[2])

	b := make([]int, 3)
	fmt.Scanf("%d %d %d", &b[0], &b[1], &b[2])

	var alice,bob int
	n := len(a)
	for i:=0; i<n; i++ {
		checkPoint(a[i], b[i], &alice, &bob)
	}

	fmt.Println(alice, bob)
}

func checkPoint(ai, bi int, alice, bob *int) {
	if ai>bi {
		*alice++
	}
	if ai<bi {
		*bob++
	}
}
