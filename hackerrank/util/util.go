package util

import (
	"os"
	"fmt"
)

//const INPUT_FILE 	= "input.txt"
const INPUT_FILE 	= "/home/namgvu/gows/src/github.com/namgivu/golang-start/hackerrank/TODO_chyymmdd_hhmm_TEMPLATE/input.txt"
const OUTPUT_FILE	= "output.txt"


func RedirectStdio2File() (inFile *os.File, outFile *os.File) {
	/*region redirect stdin to file*/

	//ref. https://stackoverflow.com/a/46399567/248616
	for { break }


	//open stream f to input file
	f, err := os.Open(INPUT_FILE)
	if err != nil { panic(err) }

	//set stdin to be f i.e. redirect stdin to input file
	os.Stdin = f

	inFile = f
	/*endregion redirect stdin to file*/

	/*region TODO redirect stdout to file*/
	outFile = nil
	/*endregion TODO redirect stdout to file*/

	return inFile, outFile
}


func ReadArray(N int) []int {
	a := make([]int, N)
	for i:=0; i<N; i++  {
		var v int
		fmt.Scanf("%d", &v)
		a[i] = v
	}

	return a
}