package day3

import (
	"testing"

	"github.com/jonathanwthom/advent-of-code-2022/helpers"
)

func TestCommonRucksackItemsSum(t *testing.T) {
	helpers.TestsForFunction(
		t,
		CommonRucksackItemsSum,
		"CommonRucksackItemsSum",
		157,
		7581,
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
