package day10

import (
	"fmt"
	"testing"

	"github.com/jonathanwthom/advent-of-code-2022/helpers"
)

func TestSignalStrengths(t *testing.T) {
	helpers.TestsForFunction(
		t,
		SignalStrengths,
		"SignalStrengths",
		13140,
		13440,
	)
}

func TestCrtImage(t *testing.T) {
	partialExp := fmt.Sprintf(`##..##..##..##..##..##..##..##..##..##..
###...###...###...###...###...###...###.
####....####....####....####....####....
#####.....#####.....#####.....#####.....
######......######......######......####
#######.......#######.......#######.....`)

	fullExp := fmt.Sprintf(`###..###..####..##..###...##..####..##..
#..#.#..#....#.#..#.#..#.#..#....#.#..#.
#..#.###....#..#....#..#.#..#...#..#..#.
###..#..#..#...#.##.###..####..#...####.
#....#..#.#....#..#.#.#..#..#.#....#..#.
#....###..####..###.#..#.#..#.####.#..#.`)

	tests := []struct {
		path        string
		expected    string
		description string
	}{
		{path: "partial_input.txt", expected: partialExp, description: "Partial input"},
		{path: "full_input.txt", expected: fullExp, description: "Full input"},
	}

	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {

			actual := CrtImage(test.path)

			if actual != test.expected {
				t.Errorf(
					"%s(%s) returned %v, expected %v",
					"CrtImage",
					test.path,
					actual,
					test.expected,
				)
			}
		})
	}
}
