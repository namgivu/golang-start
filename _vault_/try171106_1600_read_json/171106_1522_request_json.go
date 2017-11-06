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

func jsonFromGETRequest(url string) (output interface{}, err error) {
	resp, err := httpClient.Get(url)
	if err != nil {
		return nil, err
	}
	return jsonFromHttpResponse(resp)
}

func jsonFromHttpResponse(resp *http.Response) (output interface{}, err error) {
	defer resp.Body.Close()
	bodyBytes, err := ioutil.ReadAll(resp.Body)

	fmt.Printf("Response body as plain text %s\n", bodyBytes) //output json as plain text ref. https://stackoverflow.com/a/38807963/248616

	output = MyJsonArray{}
	err = json.Unmarshal(bodyBytes, &output)
	if err != nil {
		return nil, err
	}
	return output, err
	//err = json.NewDecoder(resp.Body).Decode(structOutput)
}

func main() {

	/*
	 * the url result json in the form of
		{"result":[
			{"File":"abb.yaml","Name":"ccc", ...},
			{"File":"abb.yaml","Name":"ccc", ...},
			...
		]}
	 */
	url := "http://release1.sgdc:3000/svc"

	d, err := jsonFromGETRequest(url)
	if err != nil {
		panic (fmt.Sprintf("Error occured\nDetails: %q", err))
	}
	fmt.Printf("Response body as go struct %q\n", d)
}
