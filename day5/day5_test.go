package day5

import (
	"testing"

	"github.com/jonathanwthom/advent-of-code-2022/helpers"
)

func TestTopCrates(t *testing.T) {
	tests := []struct {
		path        string
		expected    string
		description string
	}{
		{path: "partial_input.txt", expected: "CMZ", description: "Partial input"},
		{path: "full_input.txt", expected: "QNHWJVJZW", description: "Full input"},
	}

	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {

			actual := TopCrates(test.path)

			if actual != test.expected {
				t.Errorf(
					"%s(%s) returned %v, expected %v",
					"TopCrates",
					test.path,
					actual,
					test.expected,
				)
			}
		})
	}

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
