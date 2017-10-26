package main

import (
	"fmt"
)

type Base struct { //the message model used for plain/normal chat/communication
	FieldStr string
}

type Inherit struct { //struct alias
	Base
	FieldInt int
}

func main() {
	a := Base{FieldStr: "abb"}
	b := Inherit{FieldInt: 122}
	fmt.Printf("%q\n", a)
	fmt.Printf("%q\n", b)

	b.FieldStr = "ccc"
	fmt.Printf("%q\n", b)
}
