package day5

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/jonathanwthom/advent-of-code-2022/helpers"
)

// TopCrates = Part 1
func TopCrates(path string) string {
	lines := helpers.ReadLines(path)
	numStacks := getNumberofStacks(lines)
	stacks := buildStacks(lines, numStacks)

	for _, line := range lines {
		if instrRe.MatchString(line) {
			instruction := newInstruction(line)
			stacks = applyInstruction(stacks, instruction)
		}
	}

	var topCrates string
	for i := 0; i < numStacks; i++ {
		topCrates += stacks[i+1].lastCrate()
	}

	return topCrates
}

var numRe = regexp.MustCompile(`\d+`)
var letRe = regexp.MustCompile("[A-Z]")
var instrRe = regexp.MustCompile("move")

type instruction struct {
	move int
	from int
	to   int
}

func newInstruction(line string) *instruction {
	numerals := numRe.FindAllString(line, -1)
	move, _ := strconv.Atoi(numerals[0])
	from, _ := strconv.Atoi(numerals[1])
	to, _ := strconv.Atoi(numerals[2])
	return &instruction{move: move, from: from, to: to}
}

func applyInstruction(stacks map[int]*stack, instruction *instruction) map[int]*stack {
	for i := 0; i < instruction.move; i++ {
		crate := stacks[instruction.from].popCrate()
		stacks[instruction.to].pushCrate(crate)
	}

	return stacks
}

type stack struct {
	crates []string
}

// add to front of crates list, for creation only
func (s *stack) prependCrate(crate string) {
	s.crates = append([]string{crate}, s.crates...)
}

func (s *stack) pushCrate(crate string) {
	s.crates = append(s.crates, crate)
}

func (s *stack) popCrate() string {
	fmt.Println(s)
	crate := s.lastCrate()
	s.crates = s.crates[:len(s.crates)-1]
	return crate
}

func (s *stack) lastCrate() string {
	return s.crates[len(s.crates)-1]
}

func buildStacks(lines []string, numStacks int) map[int]*stack {
	stacks := map[int]*stack{}

	for i := 0; i < numStacks; i++ {
		stacks[i+1] = &stack{}
	}

	for _, line := range lines {
		if numRe.MatchString(line) {
			break
		}

		charIdx := 1
		for i := 0; i < numStacks; i++ {
			if len(line) >= charIdx+1 {
				crate := line[charIdx : charIdx+1]
				if letRe.MatchString(crate) {
					stacks[i+1].prependCrate(crate)
				}

				charIdx += 4
			}
		}
	}

	return stacks
}

func getNumberofStacks(lines []string) int {
	var num int

	for _, line := range lines {
		numerals := numRe.FindAllString(line, -1)
		if len(numerals) != 0 {
			num = len(numerals)
			break
		}
	}

	return num
}

func Part2(path string) int {
	return 0
}
