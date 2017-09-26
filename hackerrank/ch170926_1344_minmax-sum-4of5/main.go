package main

/*
https://www.hackerrank.com/challenges/mini-max-sum
*/

import (
	"fmt"
	"github.com/namgivu/golang-start/hackerrank/util"
	"math"
)

const CHALLENGE_NAME = "ch170926_1344_minmax-sum-4of5"

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

	var N int = 5
	a := make([]int64, N)
	var sum int64 = 0
	for i:=0; i<N; i++ {
		fmt.Scanf("%d", &a[i])
		sum += a[i]
	}

	var minSumo4 int64 = math.MaxInt64
	var maxSumo4 int64 = math.MinInt64
	for i:=0; i<N; i++ {
		var s int64 = sum - a[i]
		if s<minSumo4 { minSumo4 = s }
		if s>maxSumo4 { maxSumo4 = s }
	}

	fmt.Println(minSumo4, maxSumo4)

}
