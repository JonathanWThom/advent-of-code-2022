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
		453,
	)
}

func TestOverlappingPairs(t *testing.T) {
	helpers.TestsForFunction(
		t,
		OverlappingPairs,
		"OverlappingPairs",
		4,
		919,
	)
}
