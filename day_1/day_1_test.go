package day_1

import "testing"

func TestGetMostCalories(t *testing.T) {
	tests := []struct {
		path        string
		expected    int
		description string
	}{
		{path: "partial_input.txt", expected: 24000, description: "Partial input"},
		{path: "full_input.txt", expected: 71300, description: "Full input"},
	}

	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {

			actual := GetMostCalories(test.path)

			if actual != test.expected {
				t.Errorf(
					"GetMostCalories(%s) returned %v, expected %v",
					test.path,
					actual,
					test.expected,
				)
			}
		})
	}
}

func TestGetTopThreeCalories(t *testing.T) {
	tests := []struct {
		path        string
		expected    int
		description string
	}{
		{path: "partial_input.txt", expected: 45000, description: "Partial input"},
		{path: "full_input.txt", expected: 209691, description: "Full input"},
	}

	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {

			actual := GetTopThreeCalories(test.path)

			if actual != test.expected {
				t.Errorf(
					"GetTopThreeCalories(%s) returned %v, expected %v",
					test.path,
					actual,
					test.expected,
				)
			}
		})
	}
}
