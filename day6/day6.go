package day6

import (
	"strings"

	"github.com/jonathanwthom/advent-of-code-2022/helpers"
)

// StartOfPacketMarker = Part 1
func StartOfPacketMarker(path string) int {
	return getMarker(path, 4)
}

// StartOfMessageMarker = Part 2
func StartOfMessageMarker(path string) int {
	return getMarker(path, 14)
}

func getMarker(path string, markerLen int) int {
	lines := helpers.ReadLines(path)
	chars := strings.Split(lines[0], "")

loop:
	for i, char := range chars {
		if i < markerLen-1 {
			continue
		}

		recentSignals := map[string]bool{
			char: true,
		}
		for j := i - 1; j > i-markerLen; j-- {
			if recentSignals[chars[j]] == true {
				continue loop
			}
			recentSignals[chars[j]] = true
		}

		return i + 1
	}

	return 0
}
