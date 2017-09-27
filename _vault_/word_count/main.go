package main

import (
	"fmt"
	"strings"
)


func WordCount(s string) (r map[string]int) {
	r = make(map[string]int)
	splits := strings.Split(s, " ")
	for _,v := range splits {
		r[v]++
	}
	return r
}


func main() {
	s := "I am learning Go!"
	fmt.Println(WordCount(s))
}