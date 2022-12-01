package helpers

import (
	"bufio"
	"log"
	"os"
	"testing"
)

func ReadLines(path string) []string {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return lines
}

func TestsForFunction(t *testing.T, function func(string) int, funcName string, partialExp int, fullExp int) {
	tests := []struct {
		path        string
		expected    int
		description string
	}{
		{path: "partial_input.txt", expected: partialExp, description: "Partial input"},
		{path: "full_input.txt", expected: fullExp, description: "Full input"},
	}

	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {

			actual := function(test.path)

			if actual != test.expected {
				t.Errorf(
					"%s(%s) returned %v, expected %v",
					funcName,
					test.path,
					actual,
					test.expected,
				)
			}
		})
	}

}
