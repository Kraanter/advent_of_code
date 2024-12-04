package days

import (
	"adc_2024/util"
	"slices"
	"strconv"
	"strings"
)

type Day2Solution struct{}

var Day2 Day2Solution = Day2Solution{}

func (d Day2Solution) First(inputLines []string) int {
	count := 0
	for _, line := range inputLines {
		lineCorrect := true
		heights := strings.Split(line, " ")

		lastValue := 0
		for i := 0; i < len(heights)-1 && lineCorrect; i++ {
			leftNum, err := strconv.Atoi(heights[i])
			if err != nil {
				panic(1)
			}
			rightNum, err := strconv.Atoi(heights[i+1])
			if err != nil {
				panic(1)
			}
			num := leftNum - rightNum
			sign := d.getSign(num)
			if lastValue == 0 {
				lastValue = num
				continue
			}
			if sign != d.getSign(lastValue) || util.AbsoluteInt(num) > 3 || num == 0 {
				lineCorrect = false
				break
			} else {
				lastValue = num
			}
		}

		if lineCorrect {
			count++
		}
	}

	return count
}

func (d Day2Solution) Second(inputLines []string) int {
	numberListList := make([][]int, len(inputLines))
	for lineNum, line := range inputLines {
		heights := strings.Split(line, " ")
		numberList := make([]int, len(heights))

		for i, heightStr := range heights {
			heightNum, err := strconv.Atoi(heightStr)
			if err != nil {
				panic(1)
			}

			numberList[i] = heightNum
		}

		numberListList[lineNum] = numberList
	}

	count := 0
	for _, numberList := range numberListList {
		if d.checkSafety(numberList) {
			count++
			continue
		}
		for i := 0; i < len(numberList); i++ {
			listToCheck := slices.Delete(slices.Clone(numberList), i, i+1)

			if d.checkSafety(listToCheck) {
				count++
				break
			}
		}
	}

	return count
}

func (d Day2Solution) checkSafety(numbers []int) bool {
	lastDiff := 0
	for i := 0; i < len(numbers)-1; i++ {
		leftNum := numbers[i]
		rightNum := numbers[i+1]

		diff := leftNum - rightNum
		sign := d.getSign(diff)
		if lastDiff != 0 && sign != d.getSign(lastDiff) {
			return false
		}
		if util.AbsoluteInt(diff) > 3 || diff == 0 {
			return false
		}

		lastDiff = diff
	}

	return true
}

func (d Day2Solution) GetDayNumber() int {
	return 2
}

func (d Day2Solution) getSign(num int) int {
	if num == 0 {
		return 0
	}
	return num / util.AbsoluteInt(num)
}
