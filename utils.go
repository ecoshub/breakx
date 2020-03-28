package debugx

import (
	"runtime/debug"
	"strings"
)

func getLine() (string, string) {
	var (
		stackString       string
		stackNumber       int = 6
		offset            int
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
		// column char is 58
		if stackString[stackStringLength-i-1] == 58 {
			offset = stackStringLength - i - 1
			break
		}
	}
	return funcLine, stackString[offset+1:]
}
