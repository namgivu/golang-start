package main

import (
	"github.com/NodePrime/jsonpath"
	"fmt"
	"io/ioutil"
)

func main() {
	//ref. https://github.com/NodePrime/jsonpath
	//go get github.com/NodePrime/jsonpath

	jsonFile00 := "try171107_1900_json_path/test00.json"
	jsonFile01 := "try171107_1900_json_path/test01.json"
	var jsonFile = ""
	var path = ""

	jsonFile = jsonFile01 ; path="$.Items[*].title+"
	jsonFile = jsonFile01 ; path="$.Items[*].tags+"
	jsonFile = jsonFile01 ; path="$.Items[*].tags[*]+"
	jsonFile = jsonFile01 ; path="$.Items[*]?(@.title == \"A Tale of Two Cities\").tags+"

	jsonFile = jsonFile00 ; path="$.result[*].Poc+"
	jsonFile = jsonFile00 ; path="$.result[*]?(@.DiscoveryName==\"abb\").Poc+"   //CAUTION NOT working since == requires space-surrounded
	jsonFile = jsonFile00 ; path="$.result[*]?(@.DiscoveryName == \"abb\").Poc+" //this works


	paths, err := jsonpath.ParsePaths(path)
	if err!=nil {
		panic(fmt.Sprintf("%s",err))
	}

	jsonBytes, err :=  ioutil.ReadFile(jsonFile) // just pass the file name
	if err!=nil {
		panic(fmt.Sprintf("%s",err))
	}

	eval, err := jsonpath.EvalPathsInBytes(jsonBytes, paths)
	if err!=nil {
		panic(fmt.Sprintf("%s",err))
	}

	for {
		if result, ok := eval.Next(); ok {
			fmt.Println(result.Pretty(true)) // true -> show keys in pretty string
		} else {
			break
		}
	}
	if eval.Error != nil {
		panic(fmt.Sprintf("%s",eval.Error))
	}
}
