package day5

import (
	"regexp"
	"strconv"

	"github.com/jonathanwthom/advent-of-code-2022/helpers"
)

// TopCrates = Part 1
func TopCrates(path string) string {
	return getTopCrates(path, applyInstruction)
}

// CrateMover9001 = Part 2
func CrateMover9001(path string) string {
	return getTopCrates(path, applyInstruction9001)
}

func getTopCrates(path string, instructionFunc func(map[int]*stack, *instruction) map[int]*stack) string {
	lines := helpers.ReadLines(path)
	numStacks := getNumberofStacks(lines)
	stacks := buildStacks(lines, numStacks)

	for _, line := range lines {
		if instrRe.MatchString(line) {
			instruction := newInstruction(line)
			stacks = instructionFunc(stacks, instruction)
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

func applyInstruction9001(stacks map[int]*stack, instruction *instruction) map[int]*stack {
	from := stacks[instruction.from]
	to := stacks[instruction.to]
	crates := from.crates[len(from.crates)-instruction.move : len(from.crates)]
	to.crates = append(to.crates, crates...)
	from.crates = from.crates[:len(from.crates)-instruction.move]

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
