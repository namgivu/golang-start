package main

import (
	"fmt"
	"strings"
)

func main() {
	s,t := "",""
	s = "abb\n" ; t = strings.TrimSpace(s) ; fmt.Printf("%s\n",t)
	s = "abb\t" ; t = strings.TrimSpace(s) ; fmt.Printf("%s\n",t)
	s = "abb" ; s += "ccc\t" ; t = strings.TrimSpace(s) ; fmt.Printf("%s\n",t)
}
