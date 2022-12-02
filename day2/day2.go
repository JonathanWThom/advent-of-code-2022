package day2

import (
	"reflect"
	"strings"

	"github.com/jonathanwthom/advent-of-code-2022/helpers"
)

type rpsMove interface {
	vsPaper() int
	vsRock() int
	vsScissors() int
	baseVal() int
}

type rock struct{}

func (r rock) vsPaper() int    { return 0 }
func (r rock) vsRock() int     { return 3 }
func (r rock) vsScissors() int { return 6 }
func (r rock) baseVal() int    { return 1 }

type paper struct{}

func (r paper) vsPaper() int    { return 3 }
func (r paper) vsRock() int     { return 6 }
func (r paper) vsScissors() int { return 0 }
func (r paper) baseVal() int    { return 2 }

type scissors struct{}

func (r scissors) vsPaper() int    { return 6 }
func (r scissors) vsRock() int     { return 0 }
func (r scissors) vsScissors() int { return 3 }
func (r scissors) baseVal() int    { return 3 }

func getMove(rawMv string) rpsMove {
	return map[string]rpsMove{
		"X": rock{},
		"Y": paper{},
		"Z": scissors{},
		"A": rock{},
		"B": paper{},
		"C": scissors{},
	}[rawMv]
}

func getGameResult(opMv, elMv rpsMove) int {
	return map[string]int{
		"day2.rock":     elMv.vsRock(),
		"day2.paper":    elMv.vsPaper(),
		"day2.scissors": elMv.vsScissors(),
	}[reflect.TypeOf(opMv).String()]
}

func RpsScore(path string) int {
	var score int

	lines := helpers.ReadLines(path)
	for _, line := range lines {
		moves := strings.Split(line, " ")
		opMv := getMove(moves[0])
		elMv := getMove(moves[1])

		score += elMv.baseVal() + getGameResult(opMv, elMv)
	}

	return score
}

func Part2(path string) int {
	return 0
}
