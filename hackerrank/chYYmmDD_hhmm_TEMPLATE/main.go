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
		inFile, outFile := util.RedirectStdio2File()
		defer inFile.Close()
		defer outFile.Close()
	}
	/*endregion redirect stdio*/

	//read N int on line 01
	var N int
	fmt.Scanf("%d", &N)

	//read array of N integer on line 02
	a := util.ReadArray(N)

	fmt.Println(N)
	fmt.Println(a)

}
