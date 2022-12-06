package day6

import (
	"testing"

	"github.com/jonathanwthom/advent-of-code-2022/helpers"
)

func TestStartOfPacketMarker(t *testing.T) {
	helpers.TestsForFunction(
		t,
		StartOfPacketMarker,
		"StartOfPacketMarker",
		7,
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
