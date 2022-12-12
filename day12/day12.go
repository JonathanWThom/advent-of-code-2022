package day12

import (
	"fmt"
	"strings"

	"github.com/jonathanwthom/advent-of-code-2022/helpers"
)

// StepsToBestSignal = Part 1
func StepsToBestSignal(path string) int {
	squares := map[int]map[int]*square{}
	var end square
	var start square

	for y, line := range helpers.ReadLines(path) {
		squares[y] = map[int]*square{}

		for x, char := range strings.Split(line, "") {
			sq := square{y: y, x: x, char: char}

			if char == "E" {
				sq.isEnd = true
				end = sq
			}

			if char == "S" {
				sq.isStart = true
				start = sq
			}

			squares[y][x] = &sq
		}
	}

	// get all possible paths
	// find the shortest
	// somehow optimize
	// check for dead ends

	var fullPathCounts []int
	current := start
	var count int

	for {
		//fmt.Printf("current %v\n", current)
		current.visit()
		squares[current.y][current.x] = &current

		if current.equals(end) {
			fullPathCounts = append(fullPathCounts, count)
			count = 0
			break
		}

		count++

		var above *square
		aY, aX := current.aboveCoords()
		if squares[aY] != nil {
			sq := squares[aY][aX]
			above = sq
		}

		var below *square
		bY, bX := current.belowCoords()
		if squares[bY] != nil {
			sq := squares[bY][bX]
			below = sq
		}

		lY, lX := current.leftCoords()
		left := squares[lY][lX]

		rY, rX := current.rightCoords()
		right := squares[rY][rX]

		var visitable []*square
		for _, sq := range []*square{above, below, left, right} {
			if current.canVisit(sq) {
				visitable = append(visitable, sq)
			}
		}

		if len(visitable) == 0 {
			// dead end
			fmt.Println("dead end")
			current = start
			continue
		}

		var newCurrent = *visitable[0]
		for _, sq := range visitable {
			if alphaMap[sq.char] > alphaMap[newCurrent.char] {
				newCurrent = *sq
			}
		}

		current = newCurrent
	}

	return fullPathCounts[0]
}

type square struct {
	char    string
	visited bool
	isEnd   bool
	isStart bool
	x       int
	y       int
}

func (sq *square) aboveCoords() (int, int) {
	return sq.y - 1, sq.x
}

func (sq *square) belowCoords() (int, int) {
	return sq.y + 1, sq.x
}

func (sq *square) leftCoords() (int, int) {
	return sq.y, sq.x - 1
}

func (sq *square) rightCoords() (int, int) {
	return sq.y, sq.x + 1
}

func (sq *square) equals(s square) bool {
	return sq.x == s.x && sq.y == s.y
}

const alpha = "abcdefghijklmnopqrstuvwxyz"

var alphaMap = map[string]int{}

func (sq *square) canVisit(s *square) bool {
	if s == nil {
		return false
	}

	if s.visited == true {
		return false
	}

	if len(alphaMap) == 0 {
		for i, char := range strings.Split(alpha, "") {
			alphaMap[char] = i
		}
	}

	if alphaMap[s.char]-alphaMap[sq.char] > 1 {
		return false
	}

	return true
}

func (sq *square) visit() {
	sq.visited = true
}

func Part2(path string) int {
	return 0
}
