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

func TestRpsScoreWithTargetResult(t *testing.T) {
	helpers.TestsForFunction(
		t,
		RpsScoreWithTargetResult,
		"RpsScoreWithTargetResult",
		12,
		13022,
	)
}
