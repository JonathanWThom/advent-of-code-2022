package day13

import (
	"encoding/json"
	"fmt"

	"github.com/jonathanwthom/advent-of-code-2022/helpers"
)

// CorrectlyOrderedIndexSums = Part 1
func CorrectlyOrderedIndexSums(path string) int {
	pairs := getPairs(path)
	var sum int
	for _, pair := range pairs {
		if pair.inRightOrder() {
			fmt.Println(pair.id)
			sum += pair.id
		}
	}

	return sum
}

func getPairs(path string) []pair {
	pairs := []pair{}
	current := pair{id: 1}
	lines := helpers.ReadLines(path)

	id := 2
	for i, line := range helpers.ReadLines(path) {
		var pack packet
		json.Unmarshal([]byte(line), &pack)

		if current.left == nil {
			current.left = &pack
		} else if current.right == nil {
			current.right = &pack
		}

		if len(line) == 0 || i == len(lines)-1 {
			pairs = append(pairs, current)
			current = pair{id: id}
			id++
		}
	}

	return pairs
}

type packet []interface{}

type pair struct {
	id    int
	left  *packet
	right *packet
}

func inRightOrder(left, right []interface{}, id int) bool {
	if id == 2 {
		fmt.Printf("compare %v vs %v\n\n", left, right)
	}

	// wrong
	//If both values are lists, compare the first value of each list, then the second value, and so on. If the left list runs out of items first, the inputs are in the right order. If the right list runs out of items first, the inputs are not in the right order. If the lists are the same length and no comparison makes a decision about the order, continue checking the next part of the input.
	if len(right) <= len(left) {
		return false
	}

	for i := 0; i < len(left); i++ {

		leftType := "slice"
		rightType := "slice"

		switch left[i].(type) {
		case float64:
			leftType = "float64"
		}

		switch right[i].(type) {
		case float64:
			rightType = "float64"
		}

		if leftType == "float64" && rightType == "float64" {
			if left[i].(float64) <= right[i].(float64) {
				continue
			}

			return false
		}

		if leftType == "float64" && rightType == "slice" {
			if inRightOrder([]interface{}{left[i]}, right[i].([]interface{}), id) == true {
				continue
			}

			return false
		}

		if rightType == "float64" && leftType == "slice" {
			if inRightOrder(left[i].([]interface{}), []interface{}{right[i]}, id) == true {
				continue
			}

			return false
		}

		if leftType == "slice" && rightType == "slice" {
			if inRightOrder(left[i].([]interface{}), right[i].([]interface{}), id) == true {
				return false
			}
		}
	}

	return true
}

func (p *pair) inRightOrder() bool {
	left := *p.left
	right := *p.right
	return inRightOrder(left, right, p.id)
}

// Part2 = Part2
func Part2(path string) int {
	return 0
}
