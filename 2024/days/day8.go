package days

type Day8Solution struct{}

var Day8 Day8Solution = Day8Solution{}

type Vec2 struct {
	x int
	y int
}

func (d Day8Solution) First(inputLines []string) int {
	freqMap := make(map[rune][]Vec2)
	grid := make([][]rune, len(inputLines))
	for x, line := range inputLines {
		for y, char := range line {
			grid[x] = append(grid[x], '.')
			if char != '.' {
				freqMap[char] = append(freqMap[rune(char)], Vec2{x, y})
			}
		}
	}

	count := 0

	for _, list := range freqMap {
		for i, from := range list {
			for j := i + 1; j < len(list); j++ {
				to := list[j]
				dir := to.Subtract(from)

				newPos := from.Subtract(dir)
				grid, count = d.setIfPossible(grid, newPos, count)
				newPos = to.Add(dir)
				grid, count = d.setIfPossible(grid, newPos, count)
			}
		}
	}

	return count
}

func (d Day8Solution) isInBounds(grid [][]rune, location Vec2) bool {
	if location.x < 0 || location.x >= len(grid) || location.y < 0 || location.y >= len(grid[0]) {
		return false
	}

	return true
}

func (d Day8Solution) setIfPossible(grid [][]rune, location Vec2, count int) ([][]rune, int) {
	if !d.isInBounds(grid, location) {
		return grid, count
	}
	char := grid[location.x][location.y]
	if char == '.' {
		grid[location.x][location.y] = '#'
	} else if char == '#' {
		return grid, count
	}
	return grid, count + 1
}

func (d Day8Solution) countCharInGrid(grid [][]rune, charToFind rune) int {
	count := 0
	for x := 0; x < len(grid); x++ {
		for y := 0; y < len(grid[0]); y++ {
			char := grid[x][y]
			if char == charToFind {
				count++
			}
		}
	}

	return count
}

func (d Day8Solution) printGrid(grid [][]rune) {
	for x := 0; x < len(grid); x++ {
		for y := 0; y < len(grid[0]); y++ {
			char := string(grid[x][y])
			print(char)
		}
		print("\n")
	}
}

func (vecA *Vec2) Copy() Vec2 {
	return Vec2{
		x: vecA.x,
		y: vecA.y,
	}
}

func (vecA *Vec2) Add(vecB Vec2) Vec2 {
	return Vec2{
		x: vecA.x + vecB.x,
		y: vecA.y + vecB.y,
	}
}

func (vecA *Vec2) Subtract(vecB Vec2) Vec2 {
	return Vec2{
		x: vecA.x - vecB.x,
		y: vecA.y - vecB.y,
	}
}

func (d Day8Solution) Second(inputLines []string) int {
	freqMap := make(map[rune][]Vec2)
	grid := make([][]rune, len(inputLines))
	for x, line := range inputLines {
		for y, char := range line {
			grid[x] = append(grid[x], '.')
			if char != '.' {
				freqMap[char] = append(freqMap[rune(char)], Vec2{x, y})
			}
		}
	}

	count := 0

	for _, list := range freqMap {
		for i, from := range list {
			for j := i + 1; j < len(list); j++ {
				to := list[j]
				dir := to.Subtract(from)

				newPos := to.Copy()
				for d.isInBounds(grid, newPos) {
					grid, count = d.setIfPossible(grid, newPos, count)
					newPos = newPos.Add(dir)
				}
				newPos = from.Copy()
				for d.isInBounds(grid, newPos) {
					grid, count = d.setIfPossible(grid, newPos, count)
					newPos = newPos.Subtract(dir)
				}
			}
		}
	}

	return count
}

func (d Day8Solution) GetDayNumber() int {
	return 8
}
