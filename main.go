package debugx

import (
	"fmt"
	"reflect"
)

var (
	counter map[string]int = make(map[string]int, 1)
	mod     map[string]int = make(map[string]int, 1)
	enable  bool           = true
)

func BreakEvery(num int, args ...interface{}) {
	if !enable {
		return
	}
	if num <= 0 {
		fmt.Println("mod must greater than zero")
		return
	}
	funcName, lineNumber := getLine()
	name := funcName + lineNumber
	if mod[name] == 0 {
		mod[name] = num
		counter[name] = 0
		counter[name]++
		if num == 1 {
			fmt.Printf("# line:%v\tfunc:%v <Breakpoint>\n", lineNumber, funcName)
			fmt.Printf("  count: %v of %v\n", counter[name], mod[name])
		}
	} else {
		counter[name]++
		if counter[name]%mod[name] == 0 {
			if len(args) == 0 {
				fmt.Printf("# line:%v\tfunc:%v <Breakpoint>\n", lineNumber, funcName)
				fmt.Printf("  count: %v of %v\n", counter[name], mod[name])
				return
			}
			fmt.Printf("# line:%v\tfunc:%v <Breakpoint>", lineNumber, funcName)
			fmt.Printf("\n  count: %v of %v", counter[name], mod[name])
			for _, arg := range args {
				fmt.Printf("\n\t[%T]:[%v]", arg, arg)
			}
			fmt.Println()
		}
	}
}

// BreakEqual creates a breakpoint if first equal second.
func BreakEqual(first, second interface{}, args ...interface{}) {
	if !enable || !reflect.DeepEqual(first, second) {
		return
	}
	funcName, lineNumber := getLine()
	if len(args) == 0 {
		fmt.Printf("# line:%v\tfunc:%v <Breakpoint>\n", lineNumber, funcName)
		fmt.Printf("  codition: %v = %v\n", first, second)
		return
	}
	fmt.Printf("# line:%v\tfunc:%v <Breakpoint>", lineNumber, funcName)
	fmt.Printf("\n  codition: %v = %v", first, second)
	for _, arg := range args {
		fmt.Printf("\n\t[%T]:[%v]", arg, arg)
	}
	fmt.Println()
}

// BreakNotEqual creates a breakpoint if first not equal second.
func BreakNotEqual(first, second interface{}, args ...interface{}) {
	if !enable || reflect.DeepEqual(first, second) {
		return
	}
	funcName, lineNumber := getLine()
	if len(args) == 0 {
		fmt.Printf("# line:%v\tfunc:%v <Breakpoint>\n", lineNumber, funcName)
		fmt.Printf("  codition: %v != %v\n", first, second)
		return
	}
	fmt.Printf("# line:%v\tfunc:%v <Breakpoint>", lineNumber, funcName)
	fmt.Printf("\n  codition: %v != %v", first, second)
	for _, arg := range args {
		fmt.Printf("\n\t[%T]:[%v]", arg, arg)
	}
	fmt.Println()
}

// Break creates a breakpoint
// and prints its call line
// if takes any argument prints their type and values each
// with a separate line.
func Break(args ...interface{}) {
	if !enable {
		return
	}
	funcName, lineNumber := getLine()
	if len(args) == 0 {
		fmt.Printf("# line:%v\tfunc:%v <Breakpoint>\n", lineNumber, funcName)
		return
	}
	fmt.Printf("# line:%v\tfunc:%v <Breakpoint>", lineNumber, funcName)
	for _, arg := range args {
		fmt.Printf("\n\t[%T]:[%v]", arg, arg)
	}
	fmt.Println()
}

// Printif prints the value of arguments if values not nil or empty string.
func Printif(inters ...interface{}) {
	if !enable {
		return
	}
	str := ""
	for _, inter := range inters {
		stringValue := fmt.Sprint(inter)
		if stringValue != "" && stringValue != "<nil>" {
			if str != "" {
				str += " " + stringValue
			} else {
				str += stringValue
			}
		}
	}
	if str != "" {
		str += "\n"
	}
	fmt.Print(str)
}

// Enable enables prints.
func Enable() {
	enable = true
}

// Disable disables prints.
func Disable() {
	enable = false
}
