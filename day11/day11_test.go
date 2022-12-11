package day11

import (
	"testing"

	"github.com/jonathanwthom/advent-of-code-2022/helpers"
)

func TestMonkeyBusiness(t *testing.T) {
	helpers.TestsForFunction(
		t,
		MonkeyBusiness,
		"MonkeyBusiness",
		10605,
		50830,
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
