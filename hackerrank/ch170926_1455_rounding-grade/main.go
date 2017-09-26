package main

/*
https://www.hackerrank.com/challenges/grading
*/

import (
	"fmt"
	"github.com/namgivu/golang-start/hackerrank/util"
)

const CHALLENGE_NAME = "ch170926_1455_rounding-grade"

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

	for i:=0;i<N;i++ {
		var grade int
		fmt.Scanf("%d", &grade)

		/*region rounding grade*/
		diff := 5 - (grade % 5)
		var rounded int

		if diff<3 {
			rounded = grade + diff
		} else {
			rounded = grade
		}

		if grade<38 {
			rounded = grade
		}
		fmt.Println(rounded)
		/*endregion rounding grade*/

	}
}
