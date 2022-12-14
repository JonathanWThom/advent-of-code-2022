package day8

import (
	"strconv"

	"github.com/jonathanwthom/advent-of-code-2022/helpers"
)

// VisibleTrees = Part 1
func VisibleTrees(path string) int {
	var visible int

	forest := newForest(path)

	for x := 0; x < forest.height; x++ {
		for y := 0; y < forest.width; y++ {
			if x == 0 || x == forest.height-1 || y == 0 || y == forest.width-1 {
				visible++
				continue
			}

			h := forest.getValueAt(x, y)

			var invisibleCounter int

			// if any above are taller, it's not visible
			i := x - 1
			for i >= 0 {
				if forest.getValueAt(i, y) >= h {
					invisibleCounter++
					break
				}

				i--
			}

			// if any below are taller
			i = x + 1
			for i < forest.height {
				if forest.getValueAt(i, y) >= h {
					invisibleCounter++
					break
				}

				i++
			}

			// if any left are taller
			i = y - 1
			for i >= 0 {
				if forest.getValueAt(x, i) >= h {
					invisibleCounter++
					break
				}

				i--
			}

			// if any right are taller
			i = y + 1
			for i < forest.width {
				if forest.getValueAt(x, i) >= h {
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

// BestScenicScore == Part 2
func BestScenicScore(path string) int {
	forest := newForest(path)

	var bestScenicScore int

	for x := 0; x < forest.height; x++ {
		for y := 0; y < forest.width; y++ {
			h := forest.getValueAt(x, y)
			var upDist int
			var downDist int
			var leftDist int
			var rightDist int

			i := x - 1
			for i >= 0 {
				upDist++

				if forest.getValueAt(i, y) >= h {
					break
				}

				i--
			}

			i = x + 1
			for i < forest.height {
				downDist++

				if forest.getValueAt(i, y) >= h {
					break
				}

				i++
			}

			i = y - 1
			for i >= 0 {
				other := forest.getValueAt(x, i)
				if other != 0 {
					leftDist++
				}

				if other >= h {
					break
				}

				i--
			}

			i = y + 1
			for i < forest.width {
				other := forest.getValueAt(x, i)
				if other != 0 {
					rightDist++
				}

				if other >= h {
					break
				}

				i++
			}

			scenicScore := upDist * downDist * leftDist * rightDist
			if scenicScore > bestScenicScore {
				bestScenicScore = scenicScore
			}
		}
	}

	return bestScenicScore
}

type forest struct {
	lines  []string
	trees  map[int]map[int]int
	height int
	width  int
}

func newForest(path string) *forest {
	lines := helpers.ReadLines(path)

	return &forest{
		lines:  lines,
		trees:  map[int]map[int]int{},
		height: len(lines),
		width:  len(lines[0]),
	}
}

func (f *forest) getValueAt(x, y int) int {
	if f.trees[x] == nil {
		f.trees[x] = map[int]int{}
	}

	if f.trees[x][y] != 0 {
		return f.trees[x][y]
	}

	line := f.lines[x]
	height, _ := strconv.Atoi(line[y : y+1])
	f.trees[x][y] = height

	return height
}
