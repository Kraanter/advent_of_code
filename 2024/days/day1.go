package days

import (
	"adc_2024/util"
	"fmt"
	"slices"
	"strconv"
	"strings"
)

type Day1Solution struct {
	DayNr int
}

var Day1 Day1Solution = Day1Solution{}

func (d Day1Solution) First(inputLines []string) int {
	left := make([]int, len(inputLines))
	right := make([]int, len(inputLines))
	for i, line := range inputLines {
		splits := strings.Split(line, "   ")
		leftNum, err := strconv.Atoi(splits[0])
		if err != nil {
			fmt.Printf("Error while splitting left number '%v': %v", splits[0], err.Error())
			panic(1)
		}
		rightNum, err := strconv.Atoi(splits[1])
		if err != nil {
			fmt.Printf("Error while splitting right number '%v': %v", splits[1], err.Error())
			panic(1)
		}

		left[i] = leftNum
		right[i] = rightNum
	}

	slices.Sort(left)
	slices.Sort(right)

	total := 0
	for i := 0; i < len(left); i++ {
		total += util.AbsoluteInt(left[i] - right[i])
	}

	return total
}

type leftRight struct {
	left  int
	right int
}

func (d Day1Solution) Second(inputLines []string) int {
	countMap := make(map[int]*leftRight)
	addCount := func(number int, left bool) {
		if _, ok := countMap[number]; !ok {
			countMap[number] = &leftRight{
				left:  0,
				right: 0,
			}
		}

		if left {
			countMap[number].left++
		} else {
			countMap[number].right++
		}
	}

	for _, line := range inputLines {
		splits := strings.Split(line, "   ")
		leftNum, err := strconv.Atoi(splits[0])
		if err != nil {
			fmt.Printf("Error while splitting left number '%v': %v", splits[0], err.Error())
			panic(1)
		}
		rightNum, err := strconv.Atoi(splits[1])
		if err != nil {
			fmt.Printf("Error while splitting right number '%v': %v", splits[1], err.Error())
			panic(1)
		}

		addCount(leftNum, true)
		addCount(rightNum, false)
	}

	total := 0
	for k, v := range countMap {
		if v.left == 0 {
			continue
		}

		total += k * v.right * v.left
	}

	return total
}

func (d Day1Solution) GetDayNumber() int {
	return 1
}
