package breakx

import (
	"fmt"
	"runtime/debug"
	"strconv"
	"strings"
	"time"
)

const (
	stackNumber int = 3
)

func dummyStack() (path string, filename string, lineNumber int, functionName string) {
	return getLine()
}

func getLine() (path string, filename string, lineNumber int, functionName string) {
	stackString := string(debug.Stack())
	tokens := strings.Split(stackString, "\n")
	tokens = tokens[1:]
	lenTokens := len(tokens)
	paths := make([]string, 0, lenTokens/2)
	functions := make([]string, 0, lenTokens/2)
	for i := 1; i < lenTokens; i++ {
		cleanLine := strings.TrimSpace(tokens[i])
		if cleanLine == "" {
			continue
		}
		if i%2 == 0 {
			functions = append(functions, cleanLine)
		} else {
			paths = append(paths, cleanLine)
		}
	}
	pathLine := paths[stackNumber+1]
	funcLine := functions[stackNumber]
	positionOfComma := 0
	positionOfSlash := 0
	positionOfPlus := 0
	lenPathLine := len(pathLine)
	for i := 0; i < lenPathLine; i++ {
		curr := pathLine[lenPathLine-i-1]
		index := lenPathLine - i - 1
		if curr == ':' {
			positionOfComma = index
		}
		if curr == '+' {
			positionOfPlus = index
		}
		if curr == '/' {
			positionOfSlash = index
			break
		}
	}
	lineNumberString := ""
	path = pathLine[:positionOfComma]
	if positionOfPlus != 0 {
		lineNumberString = pathLine[positionOfComma+1 : positionOfPlus]
	} else {
		lineNumberString = pathLine[positionOfComma+1:]
	}
	lineNumberString = strings.TrimSpace(lineNumberString)
	lineNumber, _ = strconv.Atoi(lineNumberString)
	filename = pathLine[positionOfSlash+1 : positionOfComma]
	for i := 0; i < len(funcLine); i++ {
		if funcLine[i] == '(' {
			funcLine = funcLine[:i]
			break
		}
	}
	index := strings.Index(funcLine, ".")
	functionName = funcLine[index+1:]
	functionName += "()"
	return
}

func printCore(condname string, args ...interface{}) string {
	conditionScheme := " > %v = %v\n"
	path, filename, lineNumber, functionName := getLine()
	argc := getArgc(path, lineNumber)

	headerString := "# Breakpoint @ line:%3v, func: %v, file: %v, time: %v\n"
	header := fmt.Sprintf(headerString, lineNumber, functionName, filename, time.Now().Format(defaultTimeFormat))

	if condname != "" {
		argc = argc[:len(argc)-1]
		args = args[:len(args)-1]
	}

	if len(args) == 0 {
		return header
	}

	str := ""
	str += header
	for i, arg := range argc {
		content := fmt.Sprintf(conditionScheme, arg, args[i])
		str += content
	}
	return str
}
