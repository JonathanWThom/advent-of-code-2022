package day5

import (
	"testing"
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

func TestCrateMover9001(t *testing.T) {
	tests := []struct {
		path        string
		expected    string
		description string
	}{
		{path: "partial_input.txt", expected: "MCD", description: "Partial input"},
		{path: "full_input.txt", expected: "BPCZJLFJW", description: "Full input"},
	}

	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {

			actual := CrateMover9001(test.path)

			if actual != test.expected {
				t.Errorf(
					"%s(%s) returned %v, expected %v",
					"CrateMover9001",
					test.path,
					actual,
					test.expected,
				)
			}
		})
	}
}
