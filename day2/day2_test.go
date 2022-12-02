package day2

import (
	"testing"

	"github.com/jonathanwthom/advent-of-code-2022/helpers"
)

func TestRpsScore(t *testing.T) {
	helpers.TestsForFunction(
		t,
		RpsScore,
		"RpsScore",
		15,
		11841,
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
