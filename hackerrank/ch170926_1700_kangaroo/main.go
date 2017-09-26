package main

/*
https://www.hackerrank.com/challenges/kangaroo
*/

import (
	"fmt"
	"github.com/namgivu/golang-start/hackerrank/util"
)

const CHALLENGE_NAME = "ch170926_1700_kangaroo"

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

	var x1,v1, x2,v2 int
	fmt.Scanf("%d %d %d %d", &x1,&v1, &x2,&v2)

	const YES="YES"
	const NO="NO"

	if v1-v2!=0 {
		mod := (x1-x2) % (v1-v2)
		t   := (x1-x2) / (v1-v2) * -1
		yes := mod==0 && t>0
		var r string
		if yes { r=YES } else { r=NO }
		fmt.Println(r)
	} else {
		fmt.Println(NO)
	}
}
