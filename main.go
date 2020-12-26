package breakx

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
)

var (
	errCondition error = errors.New("Break Error: comparisson functions needs two argument")
)

var (
	enable            bool   = true
	defaultTimeFormat string = "2006-01-02 15:04:05.99"
)

// PointInfo it holds all information of Point struct
type PointInfo struct {
	LineNumber   int
	FileName     string
	FunctionName string
}

// PointStruct return PointInfo object that hold all the information of a 'Break Point'
func PointStruct() *PointInfo {
	funcName, fileName, lineNumberString := getLine()
	lineNumber, _ := strconv.Atoi(lineNumberString)
	return &PointInfo{LineNumber: lineNumber, FileName: fileName, FunctionName: funcName}
}

// Point creates a breakpoint
// and prints its call line
// if takes any argument prints their type and values each
// with a separate line.
func Point(args ...interface{}) {
	if !enable {
		return
	}
	content := printCore("", args...)
	print(content)
}

// Spoint creates a breakpoint
// and returns a string its call line
// if takes any argument prints their type and values each
// with a separate line.
func Spoint(args ...interface{}) string {
	return printCore("", args...)
}

// Pointif prints the value of arguments if values not nil or empty string.
func Pointif(inters ...interface{}) {
	if !enable {
		return
	}
	hasVal := false
	for _, inter := range inters {
		stringValue := fmt.Sprint(inter)
		if stringValue != "" && stringValue != "<nil>" {
			hasVal = true
			break
		}
	}
	if hasVal {
		content := printCore("", false)
		print(content)
	}
}

// PointEqual creates a breakpoint if first equal second.
func PointEqual(args ...interface{}) {
	if len(args) != 2 {
		fmt.Println(errCondition)
		return
	}
	if !enable || !reflect.DeepEqual(args[0], args[1]) {
		return
	}
	content := printCore("Equal", args...)
	print(content)
}

// PointNotEqual creates a breakpoint if first not equal second.
func PointNotEqual(args ...interface{}) {
	if len(args) < 2 {
		fmt.Println(errCondition)
		return
	}
	if !enable || reflect.DeepEqual(args[0], args[1]) {
		return
	}
	content := printCore("Not Equal", args...)
	print(content)
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

// PrintStatus returns print status Enable/Disable
func PrintStatus() bool {
	return enable
}

//Nop for skipping some value
func Nop(_ ...interface{}) {}
