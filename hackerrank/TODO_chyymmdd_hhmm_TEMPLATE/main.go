package main

//https://www.hackerrank.com/challenges/YOUR-TOPIC

import (
	"fmt"
	"github.com/namgivu/golang-start/hackerrank/util"
)


//turn this on when debug to redirect stdin/stdout to file in Gogland ref. TODO https://stackoverflow.com/q/46399395/248616

func main() {

	/*region redirect stdio*/
	var REDIRECT_STDIO_2_FILE bool
	REDIRECT_STDIO_2_FILE = false
	REDIRECT_STDIO_2_FILE = true

	if REDIRECT_STDIO_2_FILE {
		util.RedirectStdio2File()
	}
	/*endregion redirect stdio*/

	//TODO read N int on line 01
	//TODO read array of N integer on line 02
	//fmt.Scanf("%v\n", &N)

	fmt.Println(res)


	if REDIRECT_STDIO_2_FILE {
		//TODO ensure output buffer is flushed
	}
}
