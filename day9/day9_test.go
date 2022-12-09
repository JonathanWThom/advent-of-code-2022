package day9

import (
	"testing"

	"github.com/jonathanwthom/advent-of-code-2022/helpers"
)

func TestPart1(t *testing.T) {
	helpers.TestsForFunction(
		t,
		TailPositionsVisited,
		"TailPositionsVisited",
		13,
		6503,
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
