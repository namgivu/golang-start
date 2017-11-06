package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
	"io/ioutil"
)

type MyJsonItem map[string]string
type MyJsonArray struct {
	Result []MyJsonItem
}

var httpClient = &http.Client{Timeout: 10 * time.Second}

func getJson(url string, structOutput interface{}) (err error) {
	resp, err := httpClient.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	//output json as plain text ref. https://stackoverflow.com/a/38807963/248616
	respBodyBytes, err := ioutil.ReadAll(resp.Body)
	fmt.Printf("Response body as plain text %s\n", respBodyBytes)

	err = json.Unmarshal(respBodyBytes, structOutput)
	//err = json.NewDecoder(resp.Body).Decode(structOutput)
	return err
}

func main() {
	d := MyJsonArray{}

	/*
	 * the url result json in the form of
		{"result":[
			{"File":"abb.yaml","Name":"ccc", ...},
			{"File":"abb.yaml","Name":"ccc", ...},
			...
		]}
	 */
	url := "http://release1.sgdc:3000/svc"

	getJson(url, &d)
	fmt.Printf("Response body as go struct %q\n", d)
}
