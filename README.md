## breakx 

#### Breakx is a Simple Print Base Debugging Tool

It creates a virtual breakpoint that can act like a traditional breakpoint unless it prints the line number,function name and filename instead.

Major Functions are below.

---
#### Add a Breakpoints with `breakx.Point()`

```go
1  package main
2  
3  import "github.com/ecoshub/breakx"
4  
4  func main() {
6   // some variable for test
7   name := "eco"
8  
9   // dummy string
10  println("function started")
11 
12  // this breakpoint can read
13  // outter function name and
14  // current code line
15  breakx.Point()
16 
17  // dummy string
18  println("some other code print")
19 
20  // .. and prints arguments that given.
21  breakx.Point(name)
22 
23  // dummy string
24  println("function ended")
25 }

```
#### Output:
```
function started
# line:15  func:main.main()  <Breakpoint>
# line:21  func:main.main()  <Breakpoint>
  [string]:[eco]
some other code print
function ended
```
---
#### Conditional Breakpoints
`breakx.PointEqual()` and `breakx.PointNotEqual()`

```go
1 package main
2  
3  import "github.com/ecoshub/breakx"
4  
5  func main() {
6   // control value
7   control := 13
8  
9   // some array for iterate
10  list := []int{10, 11, 12, 13, 14, 15}
11 
12  for i, element := range list {
13    // when element is equal
14    // to control value it will print
15    // line number , function name
16    // and this case value of 'i'
17    breakx.PointEqual(element, control, i)
18  }
19 }

```
#### Output:
```
# line:17 func:main.main() <Breakpoint>
  codition: 13 = 13
  [int]:[3]
```
---
#### Counter Breakpoints with `breakx.PointEvery()`
```go
1  package main
2  
3  import (
4   "github.com/ecoshub/breakx"
5  )
6  
7  func main() {
8   // dummy loop
9   for i := 0; i <= 20; i++ {
10    // prints 'i' for every '5' iteration.
11    breakx.PointEvery(5, i)
12  }
13 }

```
#### Output:
```
# line:11 func:main.main()  <Breakpoint>
  count: 5 of 5
  [int]:[4]
# line:11 func:main.main()  <Breakpoint>
  count: 10 of 5
  [int]:[9]
# line:11 func:main.main()  <Breakpoint>
  count: 15 of 5
  [int]:[14]
# line:11 func:main.main()  <Breakpoint>
  count: 20 of 5
  [int]:[19]
```
---
#### Print if values are not nil `breakx.Printif()`
```go
package main

import (
  "github.com/ecoshub/breakx"
  "net/http"
)

func main() {
        // it will return a not-nil error
  _, err := http.Get("dummyurl")
  breakx.Printif(err)

        // it will return nil error
  _, err = http.Get("https://google.com")
  breakx.Printif(err)
}

```
#### Output:
```
Get dummyurl: unsupported protocol scheme ""
```
