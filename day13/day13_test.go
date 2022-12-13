package day13

import (
	"testing"

	"github.com/jonathanwthom/advent-of-code-2022/helpers"
)

func TestCorrectlyOrderedIndexSums(t *testing.T) {
	helpers.TestsForFunction(
		t,
		CorrectlyOrderedIndexSums,
		"CorrectlyOrderedIndexSums",
		13,
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
