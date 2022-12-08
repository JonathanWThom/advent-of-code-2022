package day8

import (
	"testing"

	"github.com/jonathanwthom/advent-of-code-2022/helpers"
)

func TestVisibleTrees(t *testing.T) {
	helpers.TestsForFunction(
		t,
		VisibleTrees,
		"VisibleTrees",
		21,
		1859,
	)
}

func TestBestScenicScore(t *testing.T) {
	helpers.TestsForFunction(
		t,
		BestScenicScore,
		"BestScenicScore",
		8,
		332640,
	)
}
