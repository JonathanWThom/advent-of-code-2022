package day10

import (
	"testing"

	"github.com/jonathanwthom/advent-of-code-2022/helpers"
)

func TestSignalStrengths(t *testing.T) {
	helpers.TestsForFunction(
		t,
		SignalStrengths,
		"SignalStrengths",
		13140,
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
