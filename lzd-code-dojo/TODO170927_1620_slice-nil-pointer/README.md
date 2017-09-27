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
