package days

import (
	"fmt"
	"strconv"
)

type Day10Solution struct{ visited []string }

var Day10 Day10Solution = Day10Solution{visited: make([]string, 0)}

func (d Day10Solution) First(inputLines []string) int {
	grid := make([][]int, len(inputLines))
	for x, line := range inputLines {
		for _, v := range line {
			num, _ := strconv.Atoi(string(v))
			grid[x] = append(grid[x], num)
		}
	}

	fmt.Printf("grid: %v\n", grid)

	total := 0
	for x, line := range grid {
		for y := range line {
			addTotal := 0
			visited := make(map[string]bool, 0)
			addTotal, _ = d.followTrailRec(grid, visited, x, y, 0)
			total += addTotal
		}
	}

	return total
}

func (d Day10Solution) followTrailRec(grid [][]int, visited map[string]bool, x, y, num int) (int, map[string]bool) {
	value, ok := d.getGridValue(grid, x, y)
	if !ok {
		return 0, visited
	}

	if _, ok := visited[d.getPosKey(x, y)]; ok {
		return 0, visited
	}

	if value == num {
		visited[d.getPosKey(x, y)] = true
		if num == 9 {
			return 1, visited
		}
		num++
		total := 0
		count := 0
		count, visited = d.followTrailRec(grid, visited, x+1, y, num)
		total += count
		count, visited = d.followTrailRec(grid, visited, x-1, y, num)
		total += count
		count, visited = d.followTrailRec(grid, visited, x, y+1, num)
		total += count
		count, visited = d.followTrailRec(grid, visited, x, y-1, num)
		total += count
		return total, visited
	}

	return 0, visited
}

func (d Day10Solution) getPosKey(x, y int) string {
	return fmt.Sprintf("%v-%v", x, y)
}

func (d Day10Solution) getGridValue(grid [][]int, x, y int) (int, bool) {
	if x < 0 || y < 0 || x >= len(grid) || y >= len(grid[0]) {
		return 0, false
	}
	return grid[x][y], true
}

func (d Day10Solution) Second(inputLines []string) int {
	grid := make([][]int, len(inputLines))
	for x, line := range inputLines {
		for _, v := range line {
			num, _ := strconv.Atoi(string(v))
			grid[x] = append(grid[x], num)
		}
	}

	total := 0
	for x, line := range grid {
		for y := range line {
			total += d.followTrailRecFull(grid, x, y, 0)
		}
	}

	return total
}

func (d Day10Solution) followTrailRecFull(grid [][]int, x, y, num int) int {
	value, ok := d.getGridValue(grid, x, y)
	if !ok {
		return 0
	}

	if value == num {
		if num == 9 {
			return 1
		}
		num++
		total := 0
		total += d.followTrailRecFull(grid, x+1, y, num)
		total += d.followTrailRecFull(grid, x-1, y, num)
		total += d.followTrailRecFull(grid, x, y+1, num)
		total += d.followTrailRecFull(grid, x, y-1, num)
		return total
	}

	return 0
}

func (d Day10Solution) GetDayNumber() int {
	return 10
}
