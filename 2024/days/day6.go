package days

import (
	"strconv"
)

type Day6Solution struct{ placedPos map[string]bool }

var Day6 Day6Solution = Day6Solution{placedPos: make(map[string]bool)}

func (d Day6Solution) First(inputLines []string) int {
	grid := make([][]rune, 0)
	var cords []int
	for x, line := range inputLines {
		gridLine := make([]rune, len(line))
		for y, v := range line {
			if v == '^' {
				cords = []int{x, y}
			}
			gridLine[y] = v
		}
		grid = append(grid, gridLine)
	}

	steps := d.traverseGrid(grid, cords, directionN)

	return steps
}

func (d Day6Solution) Second(inputLines []string) int {
	grid := make([][]rune, 0)
	var cords []int
	for x, line := range inputLines {
		gridLine := make([]rune, len(line))
		for y, v := range line {
			if v == '^' {
				cords = []int{x, y}
			}
			gridLine[y] = v
		}
		grid = append(grid, gridLine)
	}

	possibilities := d.possiblePositions(grid, cords, directionN, true)

	return possibilities
}

func (d Day6Solution) traverseGrid(grid [][]rune, cords []int, direction Direction) int {
	if d.cordsOutOfBounds(grid, cords) {
		return 0
	}

	stepCount := 0
	standingChar := grid[cords[0]][cords[1]]
	if standingChar != 'X' || standingChar == '#' {
		grid[cords[0]][cords[1]] = rune(strconv.Itoa(direction)[0])
		stepCount = 1
	}

	nextCords := d.getNextCordsInDir(cords, direction)
	for d.isObstacle(grid, nextCords) {
		direction = (direction + 2) % (directionNW + 1)
		nextCords = d.getNextCordsInDir(cords, direction)
	}

	return d.traverseGrid(grid, nextCords, direction) + stepCount
}

func (d Day6Solution) possiblePositions(grid [][]rune, cords []int, direction Direction, place bool) int {
	if d.cordsOutOfBounds(grid, cords) {
		return 0
	}

	standingChar := grid[cords[0]][cords[1]]
	if standingChar == rune(strconv.Itoa(direction)[0]) {
		return 1
	}

	// grid[cords[0]][cords[1]] = rune(strconv.Itoa(direction)[0])
	nextCords := d.getNextCordsInDir(cords, direction)
	newDirection := direction
	for d.isObstacle(grid, nextCords) {
		newDirection = (newDirection + 2) % (directionNW + 1)
		nextCords = d.getNextCordsInDir(cords, newDirection)
	}

	count := 0
	if place {
		if _, ok := d.placedPos[d.posKey(nextCords)]; !ok && !d.cordsOutOfBounds(grid, nextCords) {
			d.placedPos[d.posKey(nextCords)] = true

			gridCopy := make([][]rune, 0)
			for _, line := range grid {
				copyLine := make([]rune, len(line))
				copy(copyLine, line)
				gridCopy = append(gridCopy, copyLine)
			}

			gridCopy[nextCords[0]][nextCords[1]] = 'O'
			count = d.possiblePositions(gridCopy, cords, direction, false)
		}
	}
	grid[cords[0]][cords[1]] = rune(strconv.Itoa(direction)[0])

	return d.possiblePositions(grid, nextCords, newDirection, place) + count
}

func (d Day6Solution) posKey(cords []int) string {
	return strconv.Itoa(cords[0]) + "-" + strconv.Itoa(cords[1])
}

func (d Day6Solution) cordsOutOfBounds(grid [][]rune, cords []int) bool {
	x := cords[0]
	y := cords[1]

	if x < 0 || y < 0 || len(grid) <= x {
		return true
	}

	if len(grid) == 0 || len(grid[0]) <= y {
		return true
	}

	return false
}

func (d Day6Solution) getNextCordsInDir(cords []int, dir Direction) []int {
	x := cords[0]
	y := cords[1]
	switch dir {
	case directionN:
		return []int{x - 1, y}
	case directionE:
		return []int{x, y + 1}
	case directionS:
		return []int{x + 1, y}
	case directionW:
		return []int{x, y - 1}

	case directionNE:
		return []int{x - 1, y + 1}
	case directionNW:
		return []int{x - 1, y - 1}
	case directionSE:
		return []int{x + 1, y + 1}
	case directionSW:
		return []int{x + 1, y - 1}
	}

	return []int{-1, -1}
}

func (d Day6Solution) isObstacle(grid [][]rune, pos []int) bool {
	if d.cordsOutOfBounds(grid, pos) {
		return false
	}

	x := pos[0]
	y := pos[1]
	char := grid[x][y]

	return '#' == char || 'O' == char
}

func (d Day6Solution) printGrid(grid [][]rune) {
	for x := -1; x < len(grid); x++ {
		if x == -1 {
			print(" ")
		} else {
			print(x)
		}
		for y := 0; y < len(grid[0]); y++ {
			if x == -1 {
				print(y)
				continue
			}
			char := string(grid[x][y])
			if char == "." {
				char = " "
			} else if char == "2" || char == "6" {
				char = "-"
			} else if char == "0" || char == "4" {
				char = "|"
			}
			print(char)
		}
		print("\n")
	}
}

func (d Day6Solution) GetDayNumber() int {
	return 6
}
