package breakx

import (
	"runtime/debug"
	"strings"
)

func getLine() (string, string, string) {
	var (
		stackString       string
		line              string
		file              string
		stackNumber       int = 8
		lineOffset        int
		fileOffset        int
		stackStringLength int
		index             int
	)
	stackString = string(debug.Stack())
	tokens := strings.Split(stackString, "\n")
	tokens = tokens[1 : len(tokens)-1]
	funcLine := tokens[stackNumber]
	rootLine := tokens[stackNumber+1]
	tokens = strings.Split(rootLine, " ")
	stackString = tokens[0]
	stackStringLength = len(stackString)
	for i := 0; i < stackStringLength; i++ {
		curr := stackString[stackStringLength-i-1]
		if curr == ':' {
			lineOffset = stackStringLength - i - 1
		}
		if curr == '/' {
			fileOffset = stackStringLength - i - 1
			break
		}
	}
	line = stackString[lineOffset+1:]
	file = stackString[fileOffset+1 : lineOffset]
	for i := 0; i < len(funcLine); i++ {
		if funcLine[i] == '(' {
			funcLine = funcLine[:i]
		}
	}
	funcLine += "()"
	index = strings.Index(funcLine, ".")
	funcLine = funcLine[index+1:]
	return funcLine, file, line
}
