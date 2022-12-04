package day4

import (
	"strconv"
	"strings"

	"github.com/jonathanwthom/advent-of-code-2022/helpers"
)

// FullyContainedPairs = Part 1
func FullyContainedPairs(path string) int {
	var redundant int

	for _, line := range helpers.ReadLines(path) {
		elfAssignments := strings.Split(line, ",")
		elf1 := newElf(elfAssignments[0])
		elf2 := newElf(elfAssignments[1])

		if elf1.overlapsFullyWith(elf2) {
			redundant++
		}
	}

	return redundant
}

// OverlappingPairs = Part 2
func OverlappingPairs(path string) int {
	var overlapping int

	for _, line := range helpers.ReadLines(path) {
		elfAssignments := strings.Split(line, ",")
		elf1 := newElf(elfAssignments[0])
		elf2 := newElf(elfAssignments[1])

		if elf1.overlapsWith(elf2) {
			overlapping++
		}
	}

	return overlapping
}

type elf struct {
	start int
	end   int
}

func newElf(rawAssignment string) *elf {
	assignmentBounds := strings.Split(rawAssignment, "-")
	start, _ := strconv.Atoi(assignmentBounds[0])
	end, _ := strconv.Atoi(assignmentBounds[1])
	return &elf{start: start, end: end}
}

func (e *elf) numberOfAssignments() int {
	return e.end - e.start
}

func (e *elf) overlapsFullyWith(otherElf *elf) bool {
	moreAssignments := otherElf
	lessAssignments := e

	if e.numberOfAssignments() >= otherElf.numberOfAssignments() {
		moreAssignments = e
		lessAssignments = otherElf
	}

	return moreAssignments.start <= lessAssignments.start && moreAssignments.end >= lessAssignments.end
}

func (e *elf) overlapsWith(otherElf *elf) bool {
	startsSooner := e
	startsLater := otherElf

	if startsSooner.start > startsLater.start {
		startsSooner, startsLater = startsLater, startsSooner
	}

	return startsSooner.end >= startsLater.start
}
