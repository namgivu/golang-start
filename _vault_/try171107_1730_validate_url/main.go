package main

import (
	"fmt"
	"net"
	"time"
	"net/http"
)

func main() {
	var host, port = "",""
	host="www.google.com"
	port="80"
	fmt.Printf("%q\n", ValidateHostPort(host, port))

	fmt.Println("")

	host="www.google.com222"
	port="80"
	fmt.Printf("%q\n", ValidateHostPort(host, port))
}

func ValidateHostPort(host string, port string) (isValid bool) {
	//return ValidateHostPort11(host, port)
	//return ValidateHostPort22(host, port)
	return ValidateHostPort33(host, port)
}

func ValidateHostPort11(host string, port string) (isValid bool) {
	/*
	 * ref. https://golang.org/pkg/net/
	 * ref. https://github.com/akutz/gotil/blob/master/gotil.go#L284
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

func ValidateHostPort22(host string, port string) (isValid bool) {
	/*
	 * ref. https://stackoverflow.com/a/42227115/248616
	 */
	timeout := time.Duration(1 * time.Second)
	address := fmt.Sprintf("%s:%s", host, port)
	conn, err := net.DialTimeout("tcp",address, timeout)
	if err != nil {
		fmt.Println("Site unreachable, error: ", err)
	}

	if err != nil {
		isValid = false
		fmt.Printf("Error:%q\n", err)
	} else {
		isValid = true
		defer conn.Close()
	}
	return isValid
}

func ValidateHostPort33(host string, port string) (isValid bool) {
	/*
	 * ref. https://stackoverflow.com/a/42232301/248616
	 */
	timeout := time.Duration(1 * time.Second)
	client := http.Client{
		Timeout: timeout,
	}
	address := fmt.Sprintf("http://%s:%s", host, port)
	_, err := client.Get(address)

	if err != nil {
		isValid = false
		fmt.Printf("Error:%q\n", err)
	} else {
		isValid = true
	}
	return isValid
}
