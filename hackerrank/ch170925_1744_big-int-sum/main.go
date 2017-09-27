package main

/*
https://www.hackerrank.com/challenges/a-very-big-sum
*/

import (
	"fmt"
	"github.com/namgivu/golang-start/hackerrank/util"
	"math/big"
)

const CHALLENGE_NAME	= "ch170925_1744_big-int-sum"
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

	var N int
	fmt.Scanf("%d", &N)

	a := make([]*big.Int, N)
	for i:=0; i<N; i++ {
		var s string
		fmt.Scanf("%s", &s)

		bi := big.NewInt(0)
		bi.SetString(s, 10)

		a[i] = bi
	}

	//fmt.Println(a)

	s := big.NewInt(0)
	for _,bi := range a {
		s.Add(s,bi)
	}

	fmt.Print(s)
}
