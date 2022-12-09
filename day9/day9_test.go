package day9

import (
	"testing"

	"github.com/jonathanwthom/advent-of-code-2022/helpers"
)

func TestTailPositionsVisited(t *testing.T) {
	helpers.TestsForFunction(
		t,
		TailPositionsVisited,
		"TailPositionsVisited",
		13,
		6503,
	)
}
func TestTenKnotRope(t *testing.T) {
	t.Run("Auxillary input", func(t *testing.T) {

		actual := TenKnotRope("aux_input.txt")

		if actual != 36 {
			t.Errorf(
				"%s(%s) returned %v, expected %v",
				"TenKnotRope",
				"aux_input.txt",
				actual,
				36,
			)
		}
	})

	helpers.TestsForFunction(
		t,
		TenKnotRope,
		"TenKnotRope",
		1,
		0,
	)
}
