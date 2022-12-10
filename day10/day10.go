package day10

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/jonathanwthom/advent-of-code-2022/helpers"
)

// SignalStrengths = Part 1
func SignalStrengths(path string) int {
	x := 1
	var cycle int
	xAtCycle := map[int]int{}

	for _, line := range helpers.ReadLines(path) {
		sections := strings.Split(line, " ")
		instruction := sections[0]
		var value int
		if len(sections) > 1 {
			value, _ = strconv.Atoi(sections[1])

		}

		if instruction == "noop" {
			cycle++
			xAtCycle[cycle] = x
			continue
		}

		if instruction == "addx" {
			cycle++
			xAtCycle[cycle] = x
			cycle++
			x += value
			xAtCycle[cycle] = x
		}
	}

	var sum int
	//20th, 60th, 100th, 140th, 180th, and 220th
	for _, i := range []int{19, 59, 99, 139, 179, 219} {
		sum += xAtCycle[i] * (i + 1)
	}
	return sum
}

// CrtImage = Part 2
func CrtImage(path string) int {
	x := 1
	var cycle int
	xAtCycle := map[int]int{}

	for _, line := range helpers.ReadLines(path) {
		sections := strings.Split(line, " ")
		instruction := sections[0]
		var value int
		if len(sections) > 1 {
			value, _ = strconv.Atoi(sections[1])
		}

		if instruction == "noop" {
			cycle++
			// draw pixel
			xAtCycle[cycle] = x
			continue
		}

		if instruction == "addx" {
			cycle++
			// draw pixel
			xAtCycle[cycle] = x
			cycle++
			// draw pixel
			x += value
			xAtCycle[cycle] = x
		}
	}

	for i := 0; i < 240; i++ {
		pos := i % 40
		if xAtCycle[i] == pos || xAtCycle[i]-1 == pos || xAtCycle[i]+1 == pos {
			fmt.Printf("#")
		} else {
			fmt.Printf(".")
		}

		if pos == 0 {
			fmt.Println("")
		}

	}

	return 0
}
