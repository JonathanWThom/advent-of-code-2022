package day_1

import (
	"strconv"

	"github.com/jonathanwthom/advent-of-code-2022/helpers"
)

func GetMostCalories(path string) int {
	lines := helpers.ReadLines(path)
	var currentCals int
	var bestCals int

	for _, line := range lines {
		if line == "" {
			if currentCals > bestCals {
				bestCals = currentCals
			}
			currentCals = 0

			continue
		}

		cals, _ := strconv.Atoi(line)
		currentCals += cals
	}

	return bestCals
}
