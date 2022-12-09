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

//func TestTenKnotRope(t *testing.T) {
//helpers.TestsForFunction(
//t,
//TenKnotRope,
//"TenKnotRope",
//1,
//0,
//)
//}
