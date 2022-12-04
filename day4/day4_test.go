package day4

import (
	"testing"

	"github.com/jonathanwthom/advent-of-code-2022/helpers"
)

func TestFullyContainedPairs(t *testing.T) {
	helpers.TestsForFunction(
		t,
		FullyContainedPairs,
		"FullyContainedPairs",
		2,
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
