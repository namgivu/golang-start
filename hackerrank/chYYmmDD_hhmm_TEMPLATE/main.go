package main

/*
https://www.hackerrank.com/challenges/YOUR-TOPIC
*/

import (
	"fmt"
	"github.com/namgivu/golang-start/hackerrank/util"
)

const CHALLENGE_NAME	= "chYYmmDD_hhmm_TEMPLATE"
const ARCHIVE_PATH		= "archive"

func main() {

	/*region io setting*/

		var REDIRECT_STDIO_2_FILE bool
		REDIRECT_STDIO_2_FILE = false
		REDIRECT_STDIO_2_FILE = true //turn this on to redirect stdin/stdout to file ref. https://stackoverflow.com/q/46399395/248616

		/*region file path config*/
	var IO_HOME = util.StrFormat("{HACKERRANK_HOME}/{ARCHIVE_PATH}/{CHALLENGE_NAME}", "{HACKERRANK_HOME}", util.HACKERRANK_HOME, "{ARCHIVE_PATH}", ARCHIVE_PATH, "{CHALLENGE_NAME}", CHALLENGE_NAME)
		var INPUT_FILE			= util.StrFormat("{IO_HOME}/input.txt", "{IO_HOME}", IO_HOME)
		var OUTPUT_FILE			= util.StrFormat("{IO_HOME}/output.txt", "{IO_HOME}", IO_HOME)
		/*endregion file path config*/

		/*region redirect stdio*/
		if REDIRECT_STDIO_2_FILE {
			inFile, outFile := util.RedirectStdio2File(INPUT_FILE, OUTPUT_FILE)
			defer inFile.Close()
			defer outFile.Close()
		}
		/*endregion redirect stdio*/

	/*endregion io setting*/

	/*region solution*/

	//read N int on line 01
	var N int
	fmt.Scanf("%d", &N)

	//read array of N integer on line 02
	a := make([]int, N)
	for i:=0; i<N; i++ {
		fmt.Scanf("%d", &a[i])
	}

	fmt.Println(N)
	fmt.Println(a)

	/*endregion solution*/

}
