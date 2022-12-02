package day2

import (
	"strings"

	"github.com/jonathanwthom/advent-of-code-2022/helpers"
)

const rockNm = "rock"
const paperNm = "paper"
const scissorsNm = "scissors"

type rpsMove interface {
	vsPaper() int
	vsRock() int
	vsScissors() int
	baseVal() int
	winsAgainst() rpsMove
	drawsAgainst() rpsMove
	losesAgainst() rpsMove
	name() string
}

type rock struct{}

func (rock) vsPaper() int          { return 0 }
func (rock) vsRock() int           { return 3 }
func (rock) vsScissors() int       { return 6 }
func (rock) baseVal() int          { return 1 }
func (rock) winsAgainst() rpsMove  { return scissors{} }
func (rock) drawsAgainst() rpsMove { return rock{} }
func (rock) losesAgainst() rpsMove { return paper{} }

func (rock) name() string { return rockNm }

type paper struct{}

func (paper) vsPaper() int          { return 3 }
func (paper) vsRock() int           { return 6 }
func (paper) vsScissors() int       { return 0 }
func (paper) baseVal() int          { return 2 }
func (paper) winsAgainst() rpsMove  { return rock{} }
func (paper) drawsAgainst() rpsMove { return paper{} }
func (paper) losesAgainst() rpsMove { return scissors{} }

func (paper) name() string { return paperNm }

type scissors struct{}

func (scissors) vsPaper() int          { return 6 }
func (scissors) vsRock() int           { return 0 }
func (scissors) vsScissors() int       { return 3 }
func (scissors) baseVal() int          { return 3 }
func (scissors) winsAgainst() rpsMove  { return paper{} }
func (scissors) drawsAgainst() rpsMove { return scissors{} }
func (scissors) losesAgainst() rpsMove { return rock{} }

func (scissors) name() string { return scissorsNm }

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
		rockNm:     elMv.vsRock(),
		paperNm:    elMv.vsPaper(),
		scissorsNm: elMv.vsScissors(),
	}[opMv.name()]
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

func getMoveWithTargetResult(rawMv string, opMv rpsMove) rpsMove {
	return map[string]rpsMove{
		"X": opMv.winsAgainst(),
		"Y": opMv.drawsAgainst(),
		"Z": opMv.losesAgainst(),
	}[rawMv]
}

func RpsScoreWithTargetResult(path string) int {
	var score int

	lines := helpers.ReadLines(path)
	for _, line := range lines {
		moves := strings.Split(line, " ")
		opMv := getMove(moves[0])
		elMv := getMoveWithTargetResult(moves[1], opMv)

		score += elMv.baseVal() + getGameResult(opMv, elMv)
	}

	return score
}
