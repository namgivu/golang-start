package main

import "fmt"

/*
ref. Minh @ minh.trannguyenhao

topic
---
Input:
	An array of int64 numbers where
	1) There is exactly one number that occurs once aka. the-one number
	1) All other numbers occur exactly twice

Output:
	Find the-one number value at O(n) - don't need the index of the-one, just the value
*/

func main() {
	a := []int64 { 1,1, 2,2, 3, 4,4}
	flag := make(map[int64]int, len(a)/2+1)

	for i,v := range a {

		_,exist := flag[v]

		if exist==false { //v `v` first time => mark it as the-one candidate
			flag[v]= i

		} else if exist { //v `v` occurred before => not a candidate for sure => remove it
			delete(flag, v)
		}

	}

	//we should have 01 key only in `flag` by now
	for i,v := range flag {
		fmt.Printf("Found value=%d at a[%d]", v, i)
	}
}
