package main

import (
  "fmt"
  "strconv"
	"strings"
)

type IPAddr [4]byte


func (r IPAddr) String() string {
  s := ""
  for b := range r {
    s = s + strconv.Itoa(b) + "."
  }
  s = strings.TrimRight(s, ".")
  return s
}


func main() {
  hosts := map[string]IPAddr{
    "loopBack":  {127, 0, 0, 1},
    "googleDNS": {8, 8, 8, 8},
  }

  for name, ip := range hosts {
    fmt.Printf("%v: %v\n", name, ip)
    //fmt.Printf("%v: %v\n", name, ip.String())
  }
}
