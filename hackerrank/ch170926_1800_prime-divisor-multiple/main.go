package main

/*
https://www.hackerrank.com/challenges/between-two-sets
*/

import (
	"fmt"
	"github.com/namgivu/golang-start/hackerrank/util"
)

/*region util*/

func getMinCM(x,y int) int {
	/*maximum common divisor*/
	return x*y / getMaxCD(x,y)
}

func getMaxCD(_x,_y int) int {
	/*minimum common multiple*/

	var x,y int

	//ensure x >= y
	if x,y = _x,_y ; y>x { x,y = _y,_x }

	//ensure x,y valid
	if x==0 || y==0 { panic("Invalid value as zero for divisor") }


	if x%y==0 {
		return y
	} else {
		return getMaxCD(y, x%y)
	}
}

func countTheMiddle(y,x int) int {
	//x = y * t
	//m aka. the middle
	//m in [y, y*t1, y*t2, ... ] where t1, t2, ... is divisor of t
	mod := x % y
	if mod != 0 {
		return 0 //we cannot have the middle value
	} else { //mod MUST be zero i.e. x=y*t
		t := x / y

		//count number of divisor of t
		divisorCount := 0
		for i := 1; i <= t; i++ {
			if t%i==0 { divisorCount++ }
		}

		return divisorCount
	}
}

/*endregion util*/


const CHALLENGE_NAME = "ch170926_1800_prime-divisor-multiple"

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


	/*region load input*/
	var n,m int
	fmt.Scanf("%d %d", &n, &m)
	
	var i int
	a := make([]int, n)
	for i = 0; i < n; i++ { fmt.Scanf("%d", &a[i]) }
	b := make([]int, m)
	for i = 0; i < m; i++ { fmt.Scanf("%d", &b[i]) }
	/*endregion load input*/


	/*region the solution*/

	//find minimum common multiple aka. mcm of a array
	mcm := 1
	for i = 0; i < n; i++ {
		mcm = getMinCM(mcm, a[i])
	}

	//find maximum common divisor aka. mcd of b array
	mcd := b[0] //we have constraint m>=1 as given in the problem
	for i = 1; i < m; i++ {
		mcd = getMaxCD(mcd, b[i])
	}

	//count `the middle` values
	r := countTheMiddle(mcm,mcd)
	fmt.Println(r)

	/*endregion the solution */

}
