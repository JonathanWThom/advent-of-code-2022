package day9

import (
	"math"
	"strconv"
	"strings"

	"github.com/jonathanwthom/advent-of-code-2022/helpers"
)

// TailPositionsVisited = Part 1
func TailPositionsVisited(path string) int {
	head := knot{visited: map[int]map[int]bool{}}
	tail := knot{visited: map[int]map[int]bool{}}
	head.moveToPosition(0, 0)
	tail.moveToPosition(0, 0)

	for _, line := range helpers.ReadLines(path) {
		instructions := strings.Split(line, " ")
		direction := instructions[0]
		distance, _ := strconv.Atoi(instructions[1])

		for i := 0; i < distance; i++ {
			head.moveToPosition(xDeltas[direction], yDeltas[direction])
			//fmt.Printf("tail x: %d, tail y: %d\n", tail.x, tail.y)
			tail.follow(head)
			//fmt.Printf("head x: %d, head y: %d\n", head.x, head.y)
			//fmt.Printf("tail x: %d, tail y: %d\n", tail.x, tail.y)
			//fmt.Println("------")
		}
	}

	return tail.uniqPositions()
}

type knot struct {
	x       int
	y       int
	visited map[int]map[int]bool
}

const up = "U"
const down = "D"
const right = "R"
const left = "L"

var xDeltas = map[string]int{
	up:    1,
	down:  -1,
	right: 0,
	left:  0,
}
var yDeltas = map[string]int{
	up:    0,
	down:  0,
	right: 1,
	left:  -1,
}

func (k *knot) follow(other knot) {
	xDelta := other.x - k.x
	yDelta := other.y - k.y
	xDist := math.Abs(float64(xDelta))
	yDist := math.Abs(float64(yDelta))

	if xDist == 0 && yDist > 1 || xDist == 1 && yDist > 1 {
		if k.y > other.y {
			k.moveToPosition(xDelta, yDelta+1)
			return
		}

		k.moveToPosition(xDelta, yDelta-1)
		return
	}

	if yDist == 0 && xDist > 1 || yDist == 1 && xDist > 1 {
		if k.x > other.x {
			k.moveToPosition(xDelta+1, yDelta)
			return
		}

		k.moveToPosition(xDelta-1, yDelta)
		return
	}
}

func (k *knot) moveToPosition(xDelta, yDelta int) {
	k.x = k.x + xDelta
	k.y = k.y + yDelta

	if k.visited[k.x] == nil {
		k.visited[k.x] = map[int]bool{}
	}

	k.visited[k.x][k.y] = true
}

func (k *knot) uniqPositions() int {
	var count int
	for _, v := range k.visited {
		for range v {
			count++
		}
	}

	return count
}

func Part2(path string) int {
	return 0
}
