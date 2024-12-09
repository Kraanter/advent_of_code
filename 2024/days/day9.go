package days

import (
	"slices"
	"strconv"
	"strings"
)

type Day9Solution struct{}

var Day9 Day9Solution = Day9Solution{}

func (d Day9Solution) First(inputLines []string) int {
	line := inputLines[0]
	length := len(line)

	numbers := make([]int, length)
	for i, v := range line {
		numbers[i], _ = strconv.Atoi(string(v))
	}

	sortedNums := make([]int, 0)
	endIndex := (length / 2)
	for index, number := range numbers {
		if endIndex < 0 {
			break
		}
		if index%2 != 0 {
			skipCount := number
			var numsToAdd []int
			numsToAdd, endIndex = d.removeNFromEnd(&numbers, skipCount, endIndex)
			sortedNums = slices.Concat(sortedNums, numsToAdd)
		} else {
			identifier := index / 2
			count := number
			sortedNums = slices.Concat(sortedNums, d.listOfNValues(identifier, count))
		}
		numbers[index] = 0
	}

	total := 0
	for i, v := range sortedNums {
		total += v * i
	}

	return total
}

func (d Day9Solution) removeNFromEnd(numbers *[]int, count, endIndex int) ([]int, int) {
	numberListCopy := *numbers
	endNumIndex := endIndex * 2
	if count == 0 || endNumIndex < 0 {
		for endNumIndex > 0 && numberListCopy[endNumIndex] == 0 {
			endIndex--
			endNumIndex = endIndex * 2
		}
		return make([]int, 0), endIndex
	}
	endNum := numberListCopy[endNumIndex]
	if endNum >= count {
		numberListCopy[endNumIndex] = endNum - count
		return d.listOfNValues(endIndex, count), endIndex
	}

	thisSlice := d.listOfNValues(endIndex, endNum)
	numberListCopy[endNumIndex] = 0
	newCount := count - endNum
	alsoSlice, endIndex := d.removeNFromEnd(numbers, newCount, endIndex-1)

	return slices.Concat(thisSlice, alsoSlice), endIndex
}

func (d Day9Solution) listOfNValues(num, count int) []int {
	if count < 0 {
		panic(1)
	}
	retList := make([]int, count)
	for i := range retList {
		retList[i] = num
	}
	return retList

}

type RLEPoint struct {
	length     int
	ID         string
	numID      int
	startIndex int
	isEmpty    bool
}

func (d Day9Solution) Second(inputLines []string) int {
	line := inputLines[0]
	length := len(line)

	RLEList := make([]*RLEPoint, length)
	numberList := make([]int, 0)
	for i, v := range line {
		isEmpty := i%2 != 0
		count, _ := strconv.Atoi(string(v))
		var sliceToAdd []int
		var id string
		var numID int
		if !isEmpty {
			id = strconv.Itoa(i / 2)
			numID = i / 2
			sliceToAdd = slices.Repeat([]int{i / 2}, count)
		} else {
			id = "."
			numID = 0
			sliceToAdd = slices.Repeat([]int{0}, count)
		}
		RLEList[i] = &RLEPoint{ID: id, numID: numID, length: count, startIndex: len(numberList), isEmpty: i%2 != 0}
		numberList = slices.Concat(numberList, sliceToAdd)
	}

	for endIndex := length / 2; endIndex > 0; endIndex-- {
		pointIndex := endIndex * 2
		rlePoint := RLEList[pointIndex]
		for _, point := range RLEList[:pointIndex] {
			if point.isEmpty && point.length != 0 {
				if point.length >= rlePoint.length {
					for insertIndex := 0; insertIndex < rlePoint.length; insertIndex++ {
						numberList[point.startIndex+insertIndex] = rlePoint.numID
						numberList[rlePoint.startIndex+insertIndex] = 0
					}

					point.length -= rlePoint.length
					point.startIndex += rlePoint.length
					rlePoint.length = 0
					break
				}
			}
		}
	}

	total := 0
	for i, v := range numberList {
		total += v * i
	}

	return total
}

func (d Day9Solution) removeFullNFromEnd(numbers *[]int, count, endIndex int) ([]int, int) {
	numberListCopy := *numbers
	endNumIndex := endIndex * 2
	if count == 0 || endNumIndex < 0 {
		for endNumIndex > 0 && numberListCopy[endNumIndex] == 0 {
			endIndex--
			endNumIndex = endIndex * 2
		}
		return make([]int, 0), endIndex
	}
	endNum := numberListCopy[endNumIndex]
	if endNum >= count {
		numberListCopy[endNumIndex] = endNum - count
		return d.listOfNValues(endIndex, count), endIndex
	}

	thisSlice := d.listOfNValues(endIndex, endNum)
	numberListCopy[endNumIndex] = 0
	newCount := count - endNum
	alsoSlice, endIndex := d.removeNFromEnd(numbers, newCount, endIndex-1)

	return slices.Concat(thisSlice, alsoSlice), endIndex
}

func (d Day9Solution) swap(slicePtr *[]RLEPoint, a, b int) {
	slice := *slicePtr
	slice[a], slice[b] = slice[b], slice[a]
}

func (d Day9Solution) rleToString(point RLEPoint) string {
	return strings.Repeat(point.ID, point.length)
}

func (d Day9Solution) GetDayNumber() int {
	return 9
}
