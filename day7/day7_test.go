package day7

import (
	"testing"

	"github.com/jonathanwthom/advent-of-code-2022/helpers"
)

func TestDeletableDirectoriesSum(t *testing.T) {
	helpers.TestsForFunction(
		t,
		DeletableDirectoriesSum,
		"DeletableDirectoriesSum",
		95437,
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
