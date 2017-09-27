package main

/*
ref. https://lzd.slack.com/archives/C5XV1EJ06/p1506502834000327

topic
---
	Write a function to remove nil pointers from the slice of interfaces

	Slice must stay totally unmodified if no nil pointers found inside.

	It should work with *any* interface type, not just empty interface (`interface{}`).

sample usage
---

	```golang
	//input
	//mySliceOfAnyInterfaceType{&MyStruct{}, nil, nil, &MyStruct{}}

	removeNilPointers(&mySliceOfAnyInterfaceType)

	//output
	//mySliceOfAnyInterfaceType{&MyStruct{}, &MyStruct{}}
	```

*/

import (
	"reflect"
	"fmt"
	"errors"
)


func removeNilPointers(input interface{}) error {
	/*author: Minh @minh.trannguyenhao */

	inputValue := reflect.ValueOf(input)
	if inputValue.Kind() == reflect.Ptr { //check if `input` is a pointer; we need it to be pointer so as to remove nil element from it
		arr := inputValue.Elem() //get value of pointer `inputValue` via .Elem()

		switch arr.Kind() { //proceed with arr to be a Go slice/array only; otherwise, raise error

		case reflect.Slice:
			fallthrough
		case reflect.Array:
			i := 0
			for i < arr.Len() {
				indexAtI := arr.Index(i)
				indexAtIKind := indexAtI.Type().Kind()

				//TODO Need Minh's explanation for the below check
				okKind := false ||
					indexAtIKind == reflect.Chan ||
					indexAtIKind == reflect.Func ||
					indexAtIKind == reflect.Interface ||
					indexAtIKind == reflect.Map ||
					indexAtIKind == reflect.Ptr ||
					indexAtIKind == reflect.Slice

				if okKind && indexAtI.IsNil() {//the element at `i` "IS nil"

					//remove element at index `i` via `.AppendSlice()` instead of a normal array/slice removal
					arr = reflect.AppendSlice(arr.Slice(0, i), arr.Slice(i+1, arr.Len()))

				} else {//the element at `i` is NOT nil
					i++
				}
			}
			reflect.ValueOf(input).Elem().Set(arr)

		default:
			return errors.New("input is not pointer of array or slice")

		}

	} else {
		return errors.New("input is not pointer")
	}

	return nil
}


func main() {
	type MyStruct struct {}

	var mySliceOfAnyInterfaceType []interface{}
	mySliceOfAnyInterfaceType = []interface{} { &MyStruct{}, nil, nil, &MyStruct{} } //slice literal ref. https://tour.golang.org/moretypes/9
	fmt.Println(mySliceOfAnyInterfaceType)

	removeNilPointers(&mySliceOfAnyInterfaceType)
	fmt.Println(mySliceOfAnyInterfaceType)

}
