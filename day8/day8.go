package day8

import (
	"strconv"
	"strings"

	"github.com/jonathanwthom/advent-of-code-2022/helpers"
)

// VisibleTrees = Part 1
func VisibleTrees(path string) int {
	var visible int

	matrix := map[int]map[int]int{} // x, y, height var height int
	var width int
	var height int

	for x, line := range helpers.ReadLines(path) {
		matrix[x] = map[int]int{}
		height++
		width = len(strings.Split(line, "")) // unnecessary to do every time
		for y, rawHeight := range strings.Split(line, "") {
			height, _ := strconv.Atoi(rawHeight)
			matrix[x][y] = height
		}
	}

	// get width, get height
	// for each element in width and height, see if anything relevant is taller
	// as you do it, access or assign the value.
	// increment visible
	for x := 0; x < height; x++ {
		for y := 0; y < width; y++ {
			if x == 0 || x == height-1 || y == 0 || y == width-1 {
				visible++
				continue
			}
			h := matrix[x][y]

			var invisibleCounter int

			// if any above are taller, it's not visible
			i := x - 1
			for i >= 0 {
				if matrix[i][y] >= h {
					invisibleCounter++
					break
				}

				i--
			}

			// if any below are taller
			i = x + 1
			for i <= height {
				if matrix[i][y] >= h {
					invisibleCounter++
					break
				}

				i++
			}

			// if any left are taller
			i = y - 1
			for i >= 0 {
				if matrix[x][i] >= h {
					invisibleCounter++
					break
				}

				i--
			}

			// if any right are taller
			i = y + 1
			for i <= width {
				if matrix[x][i] >= h {
					invisibleCounter++
					break
				}

				i++
			}

			if invisibleCounter != 4 {
				visible++
			}
		}
	}

	return visible
}

func Part2(path string) int {
	return 0
}
