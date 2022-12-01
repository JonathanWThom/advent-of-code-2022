package day_1

import (
	"sort"
	"strconv"

	"github.com/jonathanwthom/advent-of-code-2022/helpers"
)

func GetTopThreeCalories(path string) int {
	counts := getSortedCalorieCounts(path)
	return counts[len(counts)-1] + counts[len(counts)-2] + counts[len(counts)-3]
}

func GetMostCalories(path string) int {
	counts := getSortedCalorieCounts(path)
	return counts[len(counts)-1]
}

func getSortedCalorieCounts(path string) []int {
	lines := helpers.ReadLines(path)
	var currentCals int
	counts := []int{}

	for i, line := range lines {
		cals, _ := strconv.Atoi(line)
		currentCals += cals

		if line == "" || (i == len(lines)-1) {
			counts = append(counts, currentCals)
			currentCals = 0
		}
	}
	sort.Ints(counts)

	return counts
}
