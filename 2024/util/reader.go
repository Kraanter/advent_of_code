package util

import (
	"fmt"
	"os"
)

func ReadInput(path string) []byte {
	bytes, err := os.ReadFile(path)
	if err != nil {
		fmt.Printf("Could not read file '%v': %v", path, err.Error())
		return make([]byte, 0)
	}

	return bytes
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
