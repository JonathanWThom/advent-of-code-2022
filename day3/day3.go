package day3

import (
	"strings"

	"github.com/jonathanwthom/advent-of-code-2022/helpers"
)

type rucksack struct {
	firstCompartment  string
	secondCompartment string
}

func (r *rucksack) getCommonItem() string {
	items := map[string]bool{}

	for _, char := range strings.Split(r.firstCompartment, "") {
		items[char] = true
	}

	for _, char := range strings.Split(r.secondCompartment, "") {
		if items[char] == true {
			return char
		}
	}

	return ""
}

const subForLowerAsciiVal = 96
const subForUpperAsciiVal = 38

func getItemValue(item string) int {
	runes := []rune(item)
	var result []int
	for i := 0; i < len(runes); i++ {
		result = append(result, int(runes[i]))
	}

	if strings.ToUpper(item) == item {
		return result[0] - subForUpperAsciiVal
	}

	return result[0] - subForLowerAsciiVal
}

func newRucksack(data string) *rucksack {
	return &rucksack{
		firstCompartment:  data[0 : len(data)/2],
		secondCompartment: data[len(data)/2:],
	}
}

func CommonRucksackItemsSum(path string) int {
	lines := helpers.ReadLines(path)
	var sum int
	for _, line := range lines {
		r := newRucksack(line)
		item := r.getCommonItem()
		sum += getItemValue(item)
	}

	return sum
}

func Part2(path string) int {
	return 0
}
