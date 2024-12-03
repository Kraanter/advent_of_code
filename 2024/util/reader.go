package util

import (
	"embed"
	"strconv"
)

var IsTestInput = false

func ReadDayInput(day int, fileSystem embed.FS) []byte {
	dayStr := strconv.Itoa(day)
	testString := ""
	if IsTestInput {
		testString = ".test"
	}
	input, err := fileSystem.ReadFile("input/day" + dayStr + testString)
	if err != nil {
		println("error: " + err.Error())
		return make([]byte, 0)
	}

	return input
}

func BytesToLines(bytes []byte) []string {
	output := make([]string, 0)
	lineNr := 0
	output = append(output, "")
	for _, byte := range bytes {
		if byte == '\n' {
			output = append(output, "")
			lineNr++
		} else {
			output[lineNr] = output[lineNr] + string(byte)
		}
	}

	return output[:len(output)-1]
}
