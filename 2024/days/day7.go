package days

import (
	"strconv"
	"strings"
)

type Day7Solution struct{}

var Day7 Day7Solution = Day7Solution{}

type DataPoint struct {
	goal  int
	parts []int
}

type Operator = int

const (
	OperatorPlus Operator = iota
	OperatorTimes
	OperatorConcat
)

func (d Day7Solution) First(inputLines []string) int {
	total := 0
	for _, line := range inputLines {
		splits := strings.Split(line, ": ")
		goal, err := strconv.Atoi(splits[0])
		if err != nil {
			panic(1)
		}
		splitParts := strings.Split(splits[1], " ")
		dataPointParts := make([]int, 0)
		for _, strPart := range splitParts {
			part, err := strconv.Atoi(strPart)
			if err != nil {
				panic(1)
			}
			dataPointParts = append(dataPointParts, part)
		}

		point := DataPoint{
			goal:  goal,
			parts: dataPointParts,
		}

		if num, ok := d.recFindOperators(point.goal, []Operator{OperatorPlus, OperatorTimes}, point.parts); ok {
			if num != len(point.parts)-1 {
				panic(1)
			}
			total += point.goal
		}
	}

	return total
}

func (d Day7Solution) Second(inputLines []string) int {
	total := 0
	for _, line := range inputLines {
		splits := strings.Split(line, ": ")
		goal, err := strconv.Atoi(splits[0])
		if err != nil {
			panic(1)
		}
		splitParts := strings.Split(splits[1], " ")
		dataPointParts := make([]int, 0)
		for _, strPart := range splitParts {
			part, err := strconv.Atoi(strPart)
			if err != nil {
				panic(1)
			}
			dataPointParts = append(dataPointParts, part)
		}

		point := DataPoint{
			goal:  goal,
			parts: dataPointParts,
		}

		if num, ok := d.recFindOperators(point.goal, []Operator{OperatorConcat, OperatorTimes, OperatorPlus}, point.parts); ok {
			if num != len(point.parts)-1 {
				panic(1)
			}
			total += point.goal
		}
	}

	return total
}

func (d Day7Solution) recFindOperators(goal int, operators []Operator, values []int) (int, bool) {
	if len(values) <= 1 {
		return 0, goal == values[0]
	}
	lastValue := values[len(values)-1]
	for _, operator := range operators {
		newGoal, ok := d.applyOppositeOperator(goal, lastValue, operator)
		if !ok {
			continue
		}
		if num, ok := d.recFindOperators(newGoal, operators, values[:len(values)-1]); ok {
			return num + 1, true
		}
	}

	return 0, false
}

func (d Day7Solution) applyOppositeOperator(a, b int, operator Operator) (int, bool) {
	switch operator {
	case OperatorPlus:
		return a - b, true
	case OperatorTimes:
		if a%b == 0 {
			return a / b, true
		}
	case OperatorConcat:
		numStr := strconv.Itoa(a)
		bStr := strconv.Itoa(b)
		if finalStr, ok := strings.CutSuffix(numStr, bStr); ok {
			num, _ := strconv.Atoi(finalStr)
			return num, true
		}
	}
	return 0, false
}

func (d Day7Solution) GetDayNumber() int {
	return 7
}
