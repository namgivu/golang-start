package main

import (
	"io/ioutil"
	"encoding/json"
	"fmt"
	"bytes"
)

/*
 * ref. https://stackoverflow.com/a/25466050/248616
 */

type YourJsonItem map[string]string
type YourJsonArray []YourJsonItem

func main() {
	jsonFile := "_vault_/try171106_1600_read_json/171106_1511_json_request.json"
	readBytes, err := ioutil.ReadFile(jsonFile)
	if err != nil {
		panic("Cannot read file")
	}

	//output json as plain text ref. https://stackoverflow.com/a/38811245/248616
	s := string(readBytes)
	fmt.Printf("As plain text %s\n", s)
	fmt.Printf("As plain text %s\n", readBytes)

	//read json as struct 01
	d1 := YourJsonArray{}
	err = json.Unmarshal(readBytes, &d1)
	fmt.Printf("As go struct %q\n", d1)

	//read json as struct 02
	d2 := YourJsonArray{}
	dec := json.NewDecoder(bytes.NewReader(readBytes))
	dec.Decode(&d2)
	fmt.Printf("As go struct %q\n", d2)
}
