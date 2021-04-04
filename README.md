# `breakx` is a functional, print base debugging tool

It creates a virtual breakpoint that can act like a traditional breakpoint

unless it can print **line number**, **current function**, **file name**, **file path**, **time** to hit the break point

and even variable **values** with variables **name**.

---

## Add a Breakpoint with `breakx.Point()`

```go
package main

import (
	"fmt"

	"github.com/ecoshub/breakx"
)

func main() {
	// some variable for test
	name := "eco"

	// dummy string
	fmt.Println("function started")

	// this breakpoint can read
	// current function name, line, file name, file path and variable names that given
	breakx.Point()

	// dummy string
	fmt.Println("some other code print")

	// .. and prints arguments that given.
	breakx.Point(name)

	// dummy string
	fmt.Println("function ended")
}


```

#### Output:

```
function started
# Breakpoint @ line: 18, func: main(), file: main.go, time: 2021-01-18 23:55:13.78
some other code print
# Breakpoint @ line: 24, func: main(), file: main.go, time: 2021-01-18 23:55:13.78
 > name = eco
function ended
```

---

## Conditional Breakpoints

`breakx.PointEqual()` and `breakx.PointNotEqual()`

```go
package main

import "github.com/ecoshub/breakx"

func main() {
	// control value
	control := 13
	// some array for iterate
	list := []int{10, 11, 12, 13, 14, 15}

	for _, element := range list {
		// when element is equal
		// to control value it will hit the breakpoint
		breakx.PointEqual(control, element)
	}
}

```

#### Output:

```
# Breakapoint @ line: 16, func: main(), file: main.go, time: 2021-01-18 23:57:01.36
 > control = 13
```

---

## Print if values are not nil `breakx.Printif()`

```go
package main

import (
	"github.com/ecoshub/breakx"
	"net/http"
)

func main() {
	// it will return a not-nil error
	// firstErr is going to trigger the break point
	_, firstErr := http.Get("https://g.qwe")

	// secondError is not going to trigger the break point
	// so will only print first errors
	_, secondError := http.Get("https://google.com")
	breakx.Printif(firstErr, secondError)
}
```

#### Output:

```
Get "https://g.qwe": dial tcp: lookup g.qwe: no such host
```

---

## Return a struct that holds the break point information `breakx.PointStruct()`

```go
package main

import (
	"github.com/ecoshub/breakx"
	"encoding/json"
	"fmt"
)

func main() {
	point := breakx.PointStruct()
	fmt.Println(point)

	// to see field names lets see point struct as json.
	j, _ := json.MarshalIndent(point, "", "  ")
	fmt.Println("point struct as json:", string(j))
}
```

#### Output:

```json
&{main.go /home/you/Desktop/test/main.go main() 10}
point struct as json: {
  "Filename": "main.go",
  "FilePath": "/home/you/Desktop/test/main.go",
  "FunctionName": "main()",
  "Line": 10
}
```

---

#### Other Functions

-	Enable or Disable all prints from breakx **breakx.Enable() / breakx.Disable()**

-	Get print status is Enabled or not **breakx.PrintStatus()**

-	Skip to use a defined variable **breakx.Nop()**
