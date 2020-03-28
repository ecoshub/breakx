package breakx

import (
	"fmt"
	"reflect"
)

var (
	counter map[string]int = make(map[string]int, 1)
	mod     map[string]int = make(map[string]int, 1)
	enable  bool           = true
)

func PointEvery(num int, args ...interface{}) {
	if !enable {
		return
	}
	funcName, lineNumber := getLine()
	name := funcName + lineNumber
	if num <= 0 && mod[name] != -1 {
		fmt.Printf("* line:%v breakx.PointEvery() function is wrong. number must greater then 0\n", lineNumber)
		mod[name] = -1
		return
	}
	switch mod[name] {
	case -1:
		return
	case 0:
		mod[name] = num
		counter[name] = 0
		counter[name]++
		if num == 1 {
			printCore(false, funcName, lineNumber, "count", "of", counter[name], mod[name], args...)
		}
	default:
		counter[name]++
		if counter[name]%mod[name] == 0 {
			printCore(false, funcName, lineNumber, "count", "of", counter[name], mod[name], args...)
		}
	}
}

// BreakEqual creates a breakpoint if first equal second.
func PointEqual(first, second interface{}, args ...interface{}) {
	if !enable || !reflect.DeepEqual(first, second) {
		return
	}
	funcName, lineNumber := getLine()
	printCore(false, funcName, lineNumber, "codition", "=", first, second, args...)
}

// BreakNotEqual creates a breakpoint if first not equal second.
func PointNotEqual(first, second interface{}, args ...interface{}) {
	if !enable || reflect.DeepEqual(first, second) {
		return
	}
	funcName, lineNumber := getLine()
	printCore(false, funcName, lineNumber, "codition", "!=", first, second, args...)
}

// Break creates a breakpoint
// and prints its call line
// if takes any argument prints their type and values each
// with a separate line.
func Point(args ...interface{}) {
	if !enable {
		return
	}
	funcName, lineNumber := getLine()
	printCore(true, funcName, lineNumber, "", "", nil, nil, args...)
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

func printCore(nocond bool, funcName, lineNumber, condname, cond string, first, second interface{}, args ...interface{}) {
	if len(args) == 0 {
		fmt.Printf("# line:%v\tfunc:%v\t<Breakpoint>\n", lineNumber, funcName)
		if !nocond {
			fmt.Printf("  %v: %v %v %v\n", condname, first, cond, second)
		}
		return
	}
	fmt.Printf("# line:%v\tfunc:%v\t<Breakpoint>", lineNumber, funcName)
	if !nocond {
		fmt.Printf("\n  %v: %v %v %v", condname, first, cond, second)
	}
	for _, arg := range args {
		fmt.Printf("\n  [%T]:[%v]", arg, arg)
	}
	fmt.Println()
}
