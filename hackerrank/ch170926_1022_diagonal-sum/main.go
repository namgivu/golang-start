package main

/*
https://www.hackerrank.com/challenges/diagonal-difference
*/

import (
	"fmt"
	"github.com/namgivu/golang-start/hackerrank/util"
	"math"
)

const CHALLENGE_NAME = "ch170926_1022_diagonal-sum"

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

	a := make([]([]int), N)
	for i:=0; i<N; i++ {
		a[i] = make([]int, N)
	}
	//fmt.Println(a)

	for i:=0; i<N; i++ {
		for j:=0; j<N; j++ {
			var aij int
			fmt.Scanf("%d", &aij)
			a[i][j] = aij
		}
	}
	//fmt.Println(a)

	sum1 := 0
	sum2 := 0
	for i:=0; i<N; i++ {
		sum1 += a[i][i] //items on diagonal 01
		sum2 += a[i][N-1 - i] //items on diagonal 02
	}
	r := math.Abs(float64(sum1-sum2))
	fmt.Print(r)
}
