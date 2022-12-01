package day_1

import (
	"testing"

	"github.com/jonathanwthom/advent-of-code-2022/helpers"
)

func TestGetMostCalories(t *testing.T) {
	helpers.TestsForFunction(
		t,
		GetMostCalories,
		"GetMostCalories",
		24000,
		71300,
	)
}

func TestGetTopThreeCalories(t *testing.T) {
	helpers.TestsForFunction(
		t,
		GetTopThreeCalories,
		"GetTopThreeCalories",
		45000,
		209691,
	)
}
