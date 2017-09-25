package main

//https://www.hackerrank.com/challenges/YOUR-TOPIC

import (
	"fmt"
	"github.com/namgivu/golang-start/hackerrank/util"
)

//const INPUT_FILE 	= "input.txt"
const INPUT_FILE	= "/home/namgvu/gows/src/github.com/namgivu/golang-start/hackerrank/chYYmmDD_hhmm_TEMPLATE/input.txt" //TODO how to use relative path here
const OUTPUT_FILE	= "/home/namgvu/gows/src/github.com/namgivu/golang-start/hackerrank/chYYmmDD_hhmm_TEMPLATE/output.txt" //TODO how to use relative path here


func main() {

	/*region redirect stdio*/
	var REDIRECT_STDIO_2_FILE bool
	REDIRECT_STDIO_2_FILE = false
	REDIRECT_STDIO_2_FILE = true //turn this on to redirect stdin/stdout to file ref. TODO https://stackoverflow.com/q/46399395/248616

	if REDIRECT_STDIO_2_FILE {
		inFile, outFile := util.RedirectStdio2File(INPUT_FILE, OUTPUT_FILE)
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
