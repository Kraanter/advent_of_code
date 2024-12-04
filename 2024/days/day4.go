package days

type Day4Solution struct{}

var Day4 Day4Solution = Day4Solution{}

type Direction = int

const (
	directionN Direction = iota
	directionNE
	directionE
	directionSE
	directionS
	directionSW
	directionW
	directionNW
)

func (d Day4Solution) First(inputLines []string) int {
	searchGrid := make([][]byte, len(inputLines))
	for lineNr, line := range inputLines {
		lineChars := make([]byte, len(line))
		for i, char := range line {
			lineChars[i] = byte(char)
		}
		searchGrid[lineNr] = lineChars
	}

	count := 0
	for x, line := range searchGrid {
		for y := range line {
			count += d.findWordCount(searchGrid, "XMAS", x, y)
		}
	}

	return count
}

func (d Day4Solution) Second(inputLines []string) int {
	searchGrid := make([][]byte, len(inputLines))
	for lineNr, line := range inputLines {
		lineChars := make([]byte, len(line))
		for i, char := range line {
			lineChars[i] = byte(char)
		}
		searchGrid[lineNr] = lineChars
	}

	count := 0
	for x, line := range searchGrid {
		for y := range line {
			if char, ok := d.getPosFromGrid(searchGrid, x, y); ok && char == 'A' {
				left := d.checkForXMASDiag(searchGrid, x, y, directionNE)
				right := d.checkForXMASDiag(searchGrid, x, y, directionSE)
				if left && right {
					count++
				}
			}
		}
	}

	return count
}

func (d Day4Solution) findWordCount(grid [][]byte, word string, xPos int, yPos int) int {
	count := 0
	char, ok := d.getPosFromGrid(grid, xPos, yPos)
	if ok && char == word[0] {
		nextWord := word[1:]
		for dir := Direction(0); dir <= directionNW; dir++ {
			xPos, yPos := d.nextPos(dir, xPos, yPos)
			if d.checkWordDirRec(grid, nextWord, xPos, yPos, dir) {
				count++
			}
		}
	}

	return count
}

func (d Day4Solution) checkWordDirRec(grid [][]byte, word string, xPos int, yPos int, direction Direction) bool {
	if len(word) == 0 {
		return true
	}

	if char, ok := d.getPosFromGrid(grid, xPos, yPos); ok && char == word[0] {
		newX, newY := d.nextPos(direction, xPos, yPos)
		return d.checkWordDirRec(grid, word[1:], newX, newY, direction)
	}

	return false
}

func (d Day4Solution) checkForXMASDiag(grid [][]byte, x, y int, dir Direction) bool {
	topX, topY := d.nextPos(dir, x, y)
	charTop, ok := d.getPosFromGrid(grid, topX, topY)
	if !ok || (charTop != 'M' && charTop != 'S') {
		return false
	}

	SWx, SWy := d.nextPos((dir+4)%(directionNW+1), x, y)
	charBot, ok := d.getPosFromGrid(grid, SWx, SWy)
	if !ok || (charBot != 'M' && charBot != 'S') || charTop == charBot {
		return false
	}

	return true
}

func (d Day4Solution) nextPos(dir Direction, x, y int) (newX, newY int) {
	switch dir {
	case directionN:
		return x + 1, y
	case directionE:
		return x, y + 1
	case directionS:
		return x - 1, y
	case directionW:
		return x, y - 1

	case directionNE:
		return x + 1, y + 1
	case directionNW:
		return x + 1, y - 1
	case directionSE:
		return x - 1, y + 1
	case directionSW:
		return x - 1, y - 1
	}

	return -1, -1
}

func (d Day4Solution) getPosFromGrid(grid [][]byte, x, y int) (byte, bool) {
	if x >= len(grid) || x < 0 || y < 0 {
		return ' ', false
	}
	line := grid[x]
	if y >= len(line) {
		return ' ', false
	}

	return line[y], true
}

func (d Day4Solution) GetDayNumber() int {
	return 4
}
