package main

import (
	"adventofcode/internal/helpers"
	"fmt"
	"slices"
	"strconv"
	"strings"
)

var year = 2024
var day = 5

func main() {
	input, _ := helpers.GetInput(year, day)

	correctlyOrderedSum, incorrectlyOrderedSum := PartOneSolution(input)

	fmt.Printf("Correct sum: %d. Incorrect sum: %d\n", correctlyOrderedSum, incorrectlyOrderedSum)
}

func PartOneSolution(input []byte) (int, int) {
	var totalCorrect int
	var totalIncorrect int

	inputParts := strings.Split(string(input), "\n\n")
	rules, orders := inputParts[0], inputParts[1]

	graph := make(map[int]map[int]bool)

	for _, rule := range strings.Split(rules, "\n") {
		parts := strings.Split(rule, "|")
		from, _ := strconv.Atoi(parts[0])
		to, _ := strconv.Atoi(parts[1])

		if graph[from] == nil {
			graph[from] = make(map[int]bool)
		}

		graph[from][to] = true
	}

	for _, order := range strings.Split(orders, "\n") {
		valid := true
		values := strings.Split(order, ",")
		var correctedOrder []int

		for _, item := range values {
			value, _ := strconv.Atoi(item)
			added := false

			for j, to := range correctedOrder {
				if graph[value][to] {
					correctedOrder = slices.Insert(correctedOrder, j, value)
					added = true
					valid = false
					break
				}
			}

			if !added {
				correctedOrder = append(correctedOrder, value)
			}
		}

		if valid {
			totalCorrect += correctedOrder[len(values)/2]
		} else {
			totalIncorrect += correctedOrder[len(values)/2]
		}
	}

	return totalCorrect, totalIncorrect
}
