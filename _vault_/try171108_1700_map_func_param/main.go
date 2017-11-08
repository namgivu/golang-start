package main

import (
	"fmt"
)

type MyMap map[string]int

func main() {
	m := make(MyMap)
	f1(m) ; fmt.Printf("%q\n", m)
	f2(&m) ; fmt.Printf("%q\n", m)
}

func f1(m MyMap) {
	m["abb"] = 122
}

func f2(m *MyMap) {
	/*We don't have pointer for map ref. https://stackoverflow.com/a/2809775/248616 - they are always reference*/
}