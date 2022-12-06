package day6

import (
	"strings"

	"github.com/jonathanwthom/advent-of-code-2022/helpers"
)

// StartOfPacketMarker = Part 1
func StartOfPacketMarker(path string) int {
	lines := helpers.ReadLines(path)
	chars := strings.Split(lines[0], "")
	values := map[int]string{}

loop:
	for i, char := range chars {
		values[i] = char
		if i < 3 {
			continue
		}

		miniValues := map[string]bool{
			char: true,
		}
		for j := i - 1; j > i-4; j-- {
			if miniValues[values[j]] == true {
				continue loop
			}
			miniValues[values[j]] = true
		}

		return i + 1
	}

	return 0
}

func Part2(path string) int {
	return 0
}
