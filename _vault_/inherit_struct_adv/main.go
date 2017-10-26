package main

import (
	"fmt"
)

type MP struct { //the message model used for plain/normal chat/communication
	FieldMP string
}

type MB struct { //the message model used for plain/normal chat/communication
	MP
}

type ME struct {
	MB
	FieldME string
}

type M struct {
	ME
}




func main() {
	a := M{ ME{FieldME: "122"} }
	fmt.Printf("%q\n", a)
}
