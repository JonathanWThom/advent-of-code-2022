package day11

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/jonathanwthom/advent-of-code-2022/helpers"
)

// MonkeyBusiness = Part 1
func MonkeyBusiness(path string) int {
	monkeys := createMonkeys(path)
	fmt.Println(monkeys[0])

	for i := 0; i < 20; i++ {
		for j := 0; j < len(monkeys); j++ {
			monkey := monkeys[j]

			fmt.Printf("Monkey %d\n", monkey.id)
			for _, item := range monkey.items {
				id, item := monkey.inspect(item)
				monkeys[id].items = append(monkeys[id].items, item)
				fmt.Printf("Item with worry level %d is thrown to monkey %d\n", item, id)
			}

			monkey.items = []int{}
		}
	}

	counts := []int{}
	for _, monkey := range monkeys {
		counts = append(counts, monkey.inspectedCount)
	}

	sort.Sort(sort.Reverse(sort.IntSlice(counts)))

	return counts[0] * counts[1]
}

func (m *monkey) inspect(item int) (int, int) {
	fmt.Printf("Monkey %d inspects an item with a worry level of %d\n", m.id, item)
	m.inspectedCount++

	worryLevel := m.worryFunc(item)
	fmt.Printf("Worry level is multiplied/added by amount to %d\n", worryLevel)
	worryLevel = worryLevel / 3 // make sure this rounds to integer
	fmt.Printf("Worry Level is divided by 3 to %d\n", worryLevel)
	if worryLevel%m.testDivisor == 0 {
		fmt.Printf("Current worry level is divisible by %d\n", m.testDivisor)
		return m.testTrueMonkey, worryLevel
	}
	fmt.Printf("Current worry level is not divisible by %d\n", m.testDivisor)
	return m.testFalseMonkey, worryLevel
}

func createMonkeys(path string) map[int]*monkey {
	monkeys := map[int]*monkey{}

	var currentMonkey *monkey

	for _, line := range helpers.ReadLines(path) {
		monkeyLine := strings.Split(line, "Monkey ")
		startingLine := strings.Split(line, "Starting items: ")
		operationLine := strings.Split(line, "Operation: ")
		testLine := strings.Split(line, "Test: ")
		ifTrueLine := strings.Split(line, "If true: ")
		ifFalseLine := strings.Split(line, "If false: ")

		if len(monkeyLine) > 1 {
			id, _ := strconv.Atoi(strings.Split(monkeyLine[1], ":")[0])
			m := &monkey{id: id}
			monkeys[id] = m
			currentMonkey = m

			continue
		}

		if len(startingLine) > 1 {
			items := strings.Split(startingLine[1], ",")
			for _, item := range items {
				itemAsInt, _ := strconv.Atoi(strings.Trim(item, " "))
				currentMonkey.items = append(currentMonkey.items, itemAsInt)
			}

			continue
		}

		if len(operationLine) > 1 {
			cmds := strings.Split(operationLine[1], " ")
			// 0 = new, 1 = equals, 2 = old, 3 = operator, // 4 = amount

			var f func(int) int

			if cmds[2] == "old" && cmds[4] == "old" {
				currentMonkey.worryFunc = func(worryVal int) int {
					return worryVal * worryVal
				}
				continue
			}

			amount, _ := strconv.Atoi(cmds[4])

			if cmds[3] == "+" {
				f = func(worryVal int) int {
					return worryVal + amount
				}
			} else {
				f = func(worryVal int) int {
					return worryVal * amount
				}
			}

			currentMonkey.worryFunc = f
			continue
		}

		if len(testLine) > 1 {
			val, _ := strconv.Atoi(strings.Split(testLine[1], " ")[2])
			currentMonkey.testDivisor = val

			continue
		}

		if len(ifTrueLine) > 1 {
			val, _ := strconv.Atoi(strings.Split(ifTrueLine[1], " ")[3])
			currentMonkey.testTrueMonkey = val

			continue
		}

		if len(ifFalseLine) > 1 {
			val, _ := strconv.Atoi(strings.Split(ifFalseLine[1], " ")[3])
			currentMonkey.testFalseMonkey = val

			continue
		}
	}

	return monkeys
}

type monkey struct {
	id              int
	inspectedCount  int
	items           []int
	worryFunc       func(int) int
	testDivisor     int
	testTrueMonkey  int
	testFalseMonkey int
}

func Part2(path string) int {
	return 0
}
