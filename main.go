package breakx

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"time"
)

var (
	errCondition error = errors.New("Break Error: comparisson functions needs two argument")
)

var (
	enable            bool   = true
	defaultTimeFormat string = "2006-01-02 15:04:05"
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
	content := printCoreBrief("", args...)
	print(content)
}

// Pointd creates a detailed breakpoint
// and prints its call line
// if takes any argument prints their type and values each
// with a separate line.
func Pointd(args ...interface{}) {
	if !enable {
		return
	}
	content := printCoreDetailed("", args...)
	print(content)
}

// Spoint creates a breakpoint
// and returns a string its call line
// if takes any argument prints their type and values each
// with a separate line.
func Spoint(args ...interface{}) string {
	return printCoreDetailed("", args...)
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
		content := printCoreBrief("", false)
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
	content := printCoreDetailed("Equal", args...)
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
	content := printCoreDetailed("Not Equal", args...)
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

func printCoreBrief(condname string, args ...interface{}) string {
	conditionScheme := "# %v = %v\t<line:%3v>\n"

	_, fileName, lineNumberString := getLine()
	lineNumber, _ := strconv.Atoi(lineNumberString)
	argc := getArgc(fileName, lineNumber)

	str := ""
	for i, arg := range args {
		content := fmt.Sprintf(conditionScheme, argc[i], arg, lineNumberString)
		str += content
	}
	return str
}

func printCoreDetailed(condname string, args ...interface{}) string {
	scheme := "# < line: %v, func: %v, file: %v, time: %v >\n"
	conditionHeader := "  \"%v\" condition triggered\n"
	conditionScheme := "\t%v (%T) = %v\n"

	funcName, fileName, lineNumberString := getLine()
	lineNumber, _ := strconv.Atoi(lineNumberString)
	argc := getArgc(fileName, lineNumber)

	if condname != "" {
		argc = argc[:len(argc)-1]
		args = args[:len(args)-1]
	}

	content := fmt.Sprintf(scheme, lineNumber, funcName, fileName, time.Now().Format(defaultTimeFormat))
	str := content
	if condname != "" {
		content := fmt.Sprintf(conditionHeader, condname)
		str += content
	}
	for i, arg := range args {
		content := fmt.Sprintf(conditionScheme, argc[i], arg, arg)
		str += content
	}
	return str
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
