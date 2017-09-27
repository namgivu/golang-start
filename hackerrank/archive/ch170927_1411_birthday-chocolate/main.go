package main

/*
https://www.hackerrank.com/challenges/the-birthday-bar
*/

import (
	"fmt"
	"github.com/namgivu/golang-start/hackerrank/util"
	"strconv"
	"strings"
	"bufio"
	"os"
)

const CHALLENGE_NAME	= "ch170927_1411_birthday-chocolate"
const ARCHIVE_PATH		= "archive"

func main() {

	/*region io setting*/

		var REDIRECT_STDIO_2_FILE bool
		REDIRECT_STDIO_2_FILE = false
		REDIRECT_STDIO_2_FILE = true //turn this on to redirect stdin/stdout to file ref. https://stackoverflow.com/q/46399395/248616

		/*region file path config*/
		var IO_HOME					= util.StrFormat("{HACKERRANK_HOME}/{CHALLENGE_NAME}", "{HACKERRANK_HOME}", util.HACKERRANK_HOME, "{CHALLENGE_NAME}", CHALLENGE_NAME)
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

	var s string
	rd := bufio.NewReader(os.Stdin)

	if s,_ = rd.ReadString('\n'); true { s = strings.Trim(s, " \n") }
	N,_ := strconv.Atoi(s)

	//read array of N integer on line 02
	if s,_ = rd.ReadString('\n'); true { s = strings.Trim(s, " \n") }
	a := make([]int, N)
	for i,v := range strings.Split(s," ") {
		a[i],_ = strconv.Atoi(v)
	}

	if s,_ = rd.ReadString('\n'); true { s = strings.Trim(s, " \n") }
	split := strings.Split(s," ")
	d,_ := strconv.Atoi(split[0])
	m,_ := strconv.Atoi(split[1])

	sum := 0
	for i:=0; i<m; i++ {
		sum += a[i]
	}

	count := 0
	if sum==d { count=1 }

	for i:=m; i<N; i++ {
		j := i-m //the index to remove from sum
		sum -= a[j]
		sum += a[i]
		if sum==d { count++ }
	}

	fmt.Println(count)

	/*endregion solution*/

}
