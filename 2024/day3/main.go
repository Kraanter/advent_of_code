package day3

import (
	"strconv"
	"strings"
)

func First(inputLines []string) int {
	total := 0
	for _, line := range inputLines {
		splits := strings.Split(line, "mul(")

		for _, curPart := range splits {
			commaSplitNums := strings.Split(curPart, ")")
			if len(commaSplitNums) == 1 && curPart == commaSplitNums[0] {
				continue
			}
			potentialNums := strings.Split(commaSplitNums[0], ",")
			if len(potentialNums) == 2 {
				num1, err := strconv.Atoi(potentialNums[0])
				if err != nil {
					continue
				}
				num2, err := strconv.Atoi(potentialNums[1])
				if err != nil {
					continue
				}

				total += num1 * num2
			}
		}
	}

	return total
}

func Second(inputLines []string) int {
	total := 0
	do := true
	doCount := 0
	dontCount := 0
	for _, line := range inputLines {
		splits := strings.Split(line, "mul(")

		for _, curPart := range splits {
			doesLoopCount := do
			doIndex := strings.LastIndex(curPart, "do()")
			dontIndex := strings.LastIndex(curPart, "don't()")

			if doIndex > dontIndex {
				doCount++
				do = true
			} else if dontIndex > doIndex {
				dontCount++
				do = false
			}

			commaSplitNums := strings.Split(curPart, ")")
			if len(commaSplitNums) == 1 && curPart == commaSplitNums[0] {
				continue
			}
			potentialNums := strings.Split(commaSplitNums[0], ",")
			if len(potentialNums) == 2 {
				num1, err := strconv.Atoi(potentialNums[0])
				if err != nil {
					continue
				}
				num2, err := strconv.Atoi(potentialNums[1])
				if err != nil {
					continue
				}

				if doesLoopCount {
					total += num1 * num2
				}
			}
		}
	}

	return total
}
