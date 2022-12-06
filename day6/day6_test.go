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
		1760,
	)
}

func TestStartOfMessageMarker(t *testing.T) {
	helpers.TestsForFunction(
		t,
		StartOfMessageMarker,
		"StartOfMessageMarker",
		19,
		2974,
	)
}
