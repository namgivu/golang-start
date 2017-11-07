package main

import (
	"fmt"
	"net"
)

func main() {
	fmt.Printf("%q\n", ValidateHostPort("www.google.com", "80"))
	fmt.Println("")
	fmt.Printf("%q\n", ValidateHostPort("www.google.com2222", "80"))
}

func ValidateHostPort(host string, port string) (isValid bool) {
	/*
	 * ref. https://golang.org/pkg/net/
	 * ref. https://github.com/akutz/gotil/blob/master/gotil.go#L284
	 * address = host:port | host | http://domain.com | etc.
	 */
	 address := fmt.Sprintf("%s:%s", host, port)
	conn, err := net.Dial("tcp", address)
	if err != nil {
		isValid = false
		fmt.Printf("Error:%q\n", err)
	} else {
		isValid = true
		defer conn.Close()
	}
	return isValid
}
