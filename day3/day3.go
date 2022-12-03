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
	return getCommonItems(r.firstCompartment, r.secondCompartment)[0]
}

func getCommonItems(str1, str2 string) []string {
	items := map[string]bool{}
	common := []string{}

	for _, char := range strings.Split(str1, "") {
		items[char] = true
	}

	for _, char := range strings.Split(str2, "") {
		if items[char] == true {
			common = append(common, char)
		}
	}

	return common
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

func RucksackBadgeSum(path string) int {
	lines := helpers.ReadLines(path)
	var sum int
	var group []string
	for _, line := range lines {
		group = append(group, line)

		if len(group) == 3 {
			firstTwoCommon := getCommonItems(group[0], group[1])
			common := getCommonItems(strings.Join(firstTwoCommon, ""), group[2])
			sum += getItemValue(common[0])
			group = []string{}
		}
	}

	return sum
}
