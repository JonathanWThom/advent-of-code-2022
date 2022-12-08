package day8

import (
	"testing"

	"github.com/jonathanwthom/advent-of-code-2022/helpers"
)

func TestVisibleTrees(t *testing.T) {
	helpers.TestsForFunction(
		t,
		VisibleTrees,
		"VisibleTrees",
		21,
		0,
	)
}

func TestPart2(t *testing.T) {
	helpers.TestsForFunction(
		t,
		Part2,
		"Part2",
		0,
		0,
	)
}
