package day10

import (
	"strconv"
	"strings"

	"github.com/jonathanwthom/advent-of-code-2022/helpers"
)

// SignalStrengths = Part 1
func SignalStrengths(path string) int {
	xAtCycle := getXAtCyclePositions(path)

	var sum int
	checkPt := 20
	for checkPt < maxCycles {
		sum += xAtCycle[checkPt-1] * checkPt
		checkPt += cycleRowLen
	}

	return sum
}

// CrtImage = Part 2
func CrtImage(path string) string {
	xAtCycle := getXAtCyclePositions(path)
	var msg string

	for i := 0; i < maxCycles; i++ {
		pos := i % cycleRowLen

		if xAtCycle[i] == pos || xAtCycle[i]-1 == pos || xAtCycle[i]+1 == pos {
			msg += lit
		} else {
			msg += dark
		}

		if pos == cycleRowLen-1 && i != maxCycles-1 {
			msg += "\n"
		}
	}

	return msg
}

const add = "addx"
const cycleRowLen = 40
const dark = "."
const lit = "#"
const maxCycles = 240
const noop = "noop"

func getXAtCyclePositions(path string) map[int]int {
	x := 1
	var cycle int
	xAtCycle := map[int]int{}

	for _, line := range helpers.ReadLines(path) {
		sections := strings.Split(line, " ")
		cmd := sections[0]
		var value int
		if len(sections) > 1 {
			value, _ = strconv.Atoi(sections[1])
		}

		if cmd == noop {
			cycle++
			xAtCycle[cycle] = x
			continue
		}

		if cmd == add {
			cycle++
			xAtCycle[cycle] = x
			cycle++
			x += value
			xAtCycle[cycle] = x
		}
	}

	return xAtCycle
}
