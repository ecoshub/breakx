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
		stackNumber       int = 6
		lineOffset        int
		fileOffset        int
		stackStringLength int
	)
	stackString = string(debug.Stack())
	tokens := strings.Split(stackString, "\n")
	// cutoff first and last
	tokens = tokens[1 : len(tokens)-1]
	funcLine := tokens[stackNumber]
	rootLine := tokens[stackNumber+1]
	tokens = strings.Split(rootLine, " ")
	// this is stack string
	stackString = tokens[0]
	stackStringLength = len(stackString)
	for i := 0; i < stackStringLength; i++ {
		curr := stackString[stackStringLength-i-1]
		// 58 is this :
		if curr == 58 {
			lineOffset = stackStringLength - i - 1
		}
		// 47 is this /
		if curr == 47 {
			fileOffset = stackStringLength - i - 1
			break
		}
	}
	line = stackString[lineOffset+1:]
	file = stackString[fileOffset+1 : lineOffset]
	for i := 0; i < len(funcLine); i++ {
		// this 40 is (
		if funcLine[i] == 40 {
			funcLine = funcLine[:i]
		}
	}
	funcLine += "()"
	return funcLine, file, line
}
