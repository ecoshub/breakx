package breakx

import (
	"io/ioutil"
	"strings"
)

func getArgc(path string, lineNumber int) []string {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		return []string{}
	}
	fileString := string(file)
	lines := strings.Split(fileString, "\n")
	line := lines[lineNumber-1]
	line = strings.TrimSpace(line)
	if line == "breakx.Point()" {
		return []string{}
	}
	startParantesis := 0
	for i := range line {
		curr := line[i]
		if curr == '(' {
			startParantesis = i
			break
		}
	}
	line = line[startParantesis-1:]
	argc := parseFunctionArguments(line)
	return argc
}

func parseFunctionArguments(args string) []string {
	argsArray := make([]string, 0, 4)
	args = strings.TrimSpace(args)
	args = strings.TrimLeft(args, "(")
	args = strings.TrimLeft(args, ")")
	inBrace := false
	start := 0
	level := 0
	for i, curr := range args {
		if curr == '(' {
			inBrace = true
			level++
			continue
		}
		if curr == ')' {
			if level == 0 {
				arg := args[start:i]
				arg = strings.TrimSpace(arg)
				argsArray = append(argsArray, arg)
				start = i + 1
			}
			inBrace = false
			level--
			continue
		}
		if inBrace {
			continue
		}
		if level == 0 {
			if curr == ',' {
				arg := args[start:i]
				arg = strings.TrimSpace(arg)
				argsArray = append(argsArray, arg)
				start = i + 1
			}
		}
	}
	return argsArray
}
