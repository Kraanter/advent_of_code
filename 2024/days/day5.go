package days

import (
	"slices"
	"strconv"
	"strings"
)

type Day5Solution struct{}

var Day5 Day5Solution = Day5Solution{}

func (d Day5Solution) First(inputLines []string) int {
	lookupMap := make(map[int][]int)
	sortedLines := make([][]int, 0)
	isSortedLines := false
	for _, line := range inputLines {
		if line == "" {
			isSortedLines = true
			continue
		}
		if isSortedLines {
			splits := strings.Split(line, ",")
			sortedLine := make([]int, len(splits))

			for i, split := range splits {
				num, err := strconv.Atoi(split)
				if err != nil {
					panic(1)
				}
				sortedLine[i] = num
			}

			sortedLines = append(sortedLines, sortedLine)
		} else {
			splits := strings.Split(line, "|")
			left, err := strconv.Atoi(splits[0])
			right, err := strconv.Atoi(splits[1])
			if err != nil {
				panic(1)
			}
			lookupMap[left] = append(lookupMap[left], right)
		}
	}

	total := 0

	for _, line := range sortedLines {
		blackList := make([]int, 0)
		correct := true
		for i := len(line) - 1; i >= 0; i-- {
			num := line[i]
			lookup, ok := lookupMap[num]
			if !ok {
				lookup = make([]int, 0)
			}
			if slices.Contains(blackList, num) {
				correct = false
				break
			}
			blackList = slices.Concat(blackList, lookup)
		}

		if correct {
			total += line[len(line)/2]
		}
	}

	return total
}

func (d Day5Solution) Second(inputLines []string) int {
	lookupMap := make(map[int]map[int]bool)
	sortedLines := make([][]int, 0)
	isSortedLines := false
	for _, line := range inputLines {
		if line == "" {
			isSortedLines = true
			continue
		}
		if isSortedLines {
			splits := strings.Split(line, ",")
			sortedLine := make([]int, len(splits))

			for i, split := range splits {
				num, err := strconv.Atoi(split)
				if err != nil {
					panic(1)
				}
				sortedLine[i] = num
			}

			sortedLines = append(sortedLines, sortedLine)
		} else {
			splits := strings.Split(line, "|")
			left, err := strconv.Atoi(splits[0])
			right, err := strconv.Atoi(splits[1])
			if err != nil {
				panic(1)
			}
			entries, ok := lookupMap[left]
			if !ok {
				lookupMap[left] = make(map[int]bool)
				entries = lookupMap[left]
			}
			entries[right] = true
		}
	}

	total := 0
	for _, line := range sortedLines {
		sortFunc := func(a, b int) int {
			if _, ok := lookupMap[a][b]; ok {
				return -1
			}
			if _, ok := lookupMap[a][b]; ok {
				return 1
			}
			return 0
		}
		if slices.IsSortedFunc(line, sortFunc) {
			continue
		}
		slices.SortFunc(line, sortFunc)

		total += line[len(line)/2]
	}

	return total
}

func (d Day5Solution) GetDayNumber() int {
	return 5
}
