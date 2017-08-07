What is so unique in Golang?
---

- 'for' loop can turn into 'while' loop
  https://tour.golang.org/flowcontrol/3
  
  ```
  for sum < 1000 {
    sum += sum
  }
  ```
  
- 'for' with range
  https://tour.golang.org/moretypes/16
  
  ```
  var pow = []int{1, 2, 4, 8}
  func main() {
    for i, v := range pow {
      fmt.Printf("index=%d value=%d\n", i, v)
    }
  }
  ```
  
  skipping index/value 
  https://tour.golang.org/moretypes/17


- 'if' condition can start with an assignment
  https://tour.golang.org/flowcontrol/6
  
  ```
  func pow() int {
    if v := 1; v < 10 {
      return v
    }
  }
  ```

- 'switch' way used for long if-else conditions
  https://tour.golang.org/flowcontrol/11
  
  ```
  func main() {
    t := time.Now()
    switch {
    case t.Hour() < 12:
      fmt.Println("Good morning!")
    case t.Hour() < 17:
      fmt.Println("Good afternoon.")
    default:
      fmt.Println("Good evening.")
    }
  }

  ```
  
- `new` the 'defer' statement
  https://tour.golang.org/flowcontrol/12
  > defers the execution of a function 'until the surrounding function returns'
  > ... 'pushed onto a stack'. When a function returns, its deferred calls are executed in 'last-in-first-out order'.
    
- also has pointer type as in C but 'no pointer arithmetic'
  https://tour.golang.org/moretypes/1
  
  pointer to a struct object is simplified i.e. p.X instead of (*p).X
  https://tour.golang.org/moretypes/4

- dynamic array in Golang is called 'slice' i.e. var s []int
  https://tour.golang.org/moretypes/7
  
- slice's len() vs cap()
  https://tour.golang.org/moretypes/11
  > len() - the number of elements it contains.
  > cap() - the number of elements in the underlying array, counting from the 1st element in the slice

- Python dict in Golang is map
  https://tour.golang.org/moretypes/19
  
  ```
  var m map[string]int
  
  func main() {
  	m = make(map[string]int)
  	m["abb"] = 1
  	fmt.Println(m["abb"])
  }
  ```
  
  test key exists
  https://tour.golang.org/moretypes/22
  ```
  elem, ok := m[key]
  ```
  
- function closure i.e. variable whose value points to a function aka. function pointer
  https://tour.golang.org/moretypes/25
  ! look quite cumbersome to me
  
- NO CLASS in Golang; instead we have 'method' defined for a received type aka. 'receiver' of the method
  https://tour.golang.org/methods/1
  ```
  type Vertex struct {
    X, Y float64
  }
  
  func (v Vertex) Abs() float64 {
    return math.Sqrt(v.X*v.X + v.Y*v.Y)
  }
  
  func main() {
    v := Vertex{3, 4}
    fmt.Println(v.Abs())
  }
  ```
  
- 'receiver' as pointer type bring back the headache of C's pointer
  https://tour.golang.org/methods/4
  
  'method' can be called from a value variable or pointer one - cause v.methodCall() ~ p.methodCall() ~ (*p).methodCall()
  https://tour.golang.org/methods/7
  
  'receiver' better to be pointer
  https://tour.golang.org/methods/8
  
- 'interface' == set of method declaration
  https://tour.golang.org/methods/9
  
  > no "implements" keyword
  https://tour.golang.org/methods/10
  
  > interface values can be thought of as a tuple of a value and a concrete type (value, type)
  https://tour.golang.org/methods/11
  
  interface's receiver can be nil and gracefully handled in Golang
  https://tour.golang.org/methods/12
  
  interface's type assert 
  https://tour.golang.org/methods/15
  ```
  t := i.(T)
  ```
  `TODO` the similarity between this syntax and that of reading from a map?
  
  popular interface is .String() i.e. ToString() in C# or str(...) in Python
  https://tour.golang.org/methods/17
  
  standard Reader interface
  https://tour.golang.org/methods/21
  ```
  func (T) Read(b []byte) (n int, err error)
  ```
  
- interface review
  - interface's implementation could appear in any package without pre-arrangement
  - method = a function injected into a specific type
  - `NO` interface = {method01, method02, ...}
  - `YES` interface = {function prototype 01, function prototype 02, ...}

  - golang object = { fields, methods}
    + fields  defined by 'struct'
    + methods defined by 'interface'
    + interface implemented by 'function's type-injection'
    