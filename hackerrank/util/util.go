package util

import (
	"os"
	"fmt"
	"strings"
)


func RedirectStdio2File(INPUT_FILE, OUTPUT_FILE string) (inFile, outFile *os.File) {

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


	/*region redirect stdout to file*/

		//open stream f2 to output file
		f2, err := os.Create(OUTPUT_FILE)
		if err != nil { panic(err) }

		//set stdout to be f2 i.e. redirect stdin to input file
		os.Stdout = f2

		outFile = f2
	/*endregion redirect stdout to file*/


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


func StrFormat(format string, args ...string) string {
	/*ref. https://stackoverflow.com/a/40811635/248616*/
	r := strings.NewReplacer(args...)
	return r.Replace(format)
}
