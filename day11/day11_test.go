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

func TestMonkeyBusinessTenThousand(t *testing.T) {
	helpers.TestsForFunction(
		t,
		MonkeyBusinessTenThousand,
		"MonkeyBusinessTenThousand",
		2713310158,
		14399640002,
	)
}
