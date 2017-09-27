package main

/*
https://www.hackerrank.com/challenges/apple-and-orange
*/

import (
	"fmt"
	"github.com/namgivu/golang-start/hackerrank/util"
)

const CHALLENGE_NAME	= "ch170926_1500_count-apple-orange"
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

	var s,t,a,b,m,n int
	fmt.Scanf("%d %d\n%d %d\n%d %d", &s,&t, &a,&b, &m,&n)

	apples := make([]int, m)
	for i:=0; i<m; i++ { fmt.Scanf("%d", &apples[i]) }
	oranges := make([]int, n)
	for i:=0; i<n; i++ { fmt.Scanf("%d", &oranges[i]) }

	countApples := 0
	for i := 0; i < m; i++ {
		d := apples[i]
		pos := a + d
		if pos>=s && pos<=t {
			countApples++
		}
	}

	countOranges := 0
	for i := 0; i < n; i++ {
		d := oranges[i]
		pos := b + d
		if pos>=s && pos<=t {
			countOranges++
		}
	}

	fmt.Println(countApples)
	fmt.Println(countOranges)
}
