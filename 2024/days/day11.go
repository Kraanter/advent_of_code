package days

import (
	"slices"
	"strconv"
	"strings"
)

type Day11Solution struct{}

var Day11 Day11Solution = Day11Solution{}

func (d Day11Solution) First(inputLines []string) int {
	strStones := strings.Split(inputLines[0], " ")
	stones := make([]int, len(strStones))
	for i, v := range strStones {
		stones[i], _ = strconv.Atoi(v)
	}

	for i := 0; i < 25; i++ {
		stonesCopy := make([]int, len(stones))
		copy(stonesCopy, stones)
		stones = make([]int, 0)
		for _, stone := range stonesCopy {
			newStones := d.applyRules(stone)
			stones = slices.Concat(stones, newStones)
		}
	}

	return len(stones)
}

func (d Day11Solution) applyRules(stone int) []int {
	switch {
	case stone == 0:
		return []int{1}
	case d.digitCount(stone)%2 == 0:
		return d.rule2(stone)
	default:
		return []int{stone * 2024}
	}
}

func (d Day11Solution) rule2(num int) []int {
	numStr := strconv.Itoa(num)
	halfPoint := len(numStr) / 2
	left, right := numStr[:halfPoint], numStr[halfPoint:]
	leftNum, _ := strconv.Atoi(left)
	rightNum, _ := strconv.Atoi(right)

	return []int{leftNum, rightNum}
}

func (d Day11Solution) digitCount(num int) int {
	return len(strconv.Itoa(num))
}

// Poging to linkedlist, werkt niet want toevoegen zorgt ervoor dat je die gelijk ook weer doorloopt dus werkt niet zoals gewild
// Oplossing hiervoor is het kopieren van de linkedlist maar dan ben je hetzelfde aant doen als bij first
//	type Node struct {
//		stone int
//		next  *Node
//	}
//
//	type LinkedList struct {
//		first *Node
//		last  *Node
//	}
//
//	func (l *LinkedList) Insert(value int) {
//		if l.first == nil {
//			l.first = &Node{stone: value}
//			l.last = l.first
//		} else {
//			l.last.next = &Node{stone: value}
//		}
//	}
//
//	func (l *LinkedList) Length() int {
//		node := l.first
//		count := 0
//		for node != nil {
//			count++
//			node = node.next
//		}
//		return count
//	}
//
//	func (l *LinkedList) Print() {
//		node := l.first
//		count := 0
//		for node != nil {
//			count++
//			println(count, node.stone)
//			node = node.next
//		}
//	}
//
//	func (n *Node) InsretInAndAfter(nums []int) {
//		n.stone = nums[0]
//		next := n.next
//		if len(nums) > 1 {
//			for i := len(nums) - 1; i > 0; i-- {
//				newNode := &Node{stone: nums[i], next: next}
//				next = newNode
//			}
//		}
//		n.next = next
//	}

func (d Day11Solution) Second(inputLines []string) int {
	strStones := strings.Split(inputLines[0], " ")
	stones := make(map[int]int)
	for _, v := range strStones {
		stone, _ := strconv.Atoi(v)
		stones[stone]++
	}

	for i := 0; i < 75; i++ {
		newStones := make(map[int]int)

		for stone, count := range stones {
			newStonesToAdd := d.applyRules(stone)
			for _, stone := range newStonesToAdd {
				newStones[stone] += count
			}
		}
		stones = newStones
	}

	total := 0
	for _, count := range stones {
		total += count
	}

	return total
}

func (d Day11Solution) GetDayNumber() int {
	return 11
}
