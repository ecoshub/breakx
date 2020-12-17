package breakx

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

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

func getArgc(fileName string, lineNumber int) []string {
	workingDir, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		return nil
	}
	workingDir = workingDir + "/" + fileName
	file, err := ioutil.ReadFile(workingDir)
	count := 0
	lastNewLine := 0
	newLine := 0
	for i, chr := range file {
		if chr == '\n' {
			if count == lineNumber-1 {
				newLine = i
				break
			} else {
				lastNewLine = i
			}
			count++
		}
	}
	line := string(file[lastNewLine:newLine])
	line = strings.TrimSpace(line)

	braceIndex := strings.Index(line, "(")
	line = line[braceIndex:]
	return parseFunctionArguments(line)
}
