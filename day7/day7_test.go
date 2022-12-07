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
		1582412,
	)
}

func TestDeletableDirectorySize(t *testing.T) {
	helpers.TestsForFunction(
		t,
		DeletableDirectorySize,
		"DeletableDirectorySize",
		24933642,
		3696336,
	)
}
