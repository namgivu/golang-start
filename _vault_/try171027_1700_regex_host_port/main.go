package main

import (
	"fmt"
	"strings"
	"regexp"
)

func main() {
	address := "1.2.3.4:55"
	host, port, ok := checkHostPort(address)
	fmt.Printf("%s %s %v", host, port, ok)
}


/*region checkHostPort*/
const REGEX_NUMBER = "[0-9]+"
const REGEX_IP = "^(([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])$"
const REGEX_HOST_NAME = "^(([a-zA-Z0-9]|[a-zA-Z0-9][a-zA-Z0-9\\-]*[a-zA-Z0-9])\\.)*([A-Za-z0-9]|[A-Za-z0-9][A-Za-z0-9\\-]*[A-Za-z0-9])$"

func checkHostPort(s string) (host string, port string, isIpPort bool) {
	//ensure we have ip:host format
	sp := strings.Split(s, ":")
	if len(sp)!=2 { return "","",false }
	host, port = sp[0], sp[1]

	//ensure valid ip address or a host name ref. https://stackoverflow.com/a/106223/248616
	if !isRegexMatched(host, REGEX_IP) && !isRegexMatched(host, REGEX_HOST_NAME) { return "","",false }

	//ensure port is number ref. https://stackoverflow.com/a/32987402/248616
	if !isRegexMatched(port, REGEX_NUMBER) { return "","",false }

	return host, port, true
}

func isRegexMatched(port string, regex string) (matched bool) {
	/* ref. https://golang.org/pkg/regexp/#example_ */
	re := regexp.MustCompile(regex)
	matched = re.MatchString(port)
	return matched
}
/*endregion checkHostPort*/
