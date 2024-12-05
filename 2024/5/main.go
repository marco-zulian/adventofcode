package main

import (
	"adventofcode/internal/helpers"
	"fmt"
	"strconv"
	"strings"
)

var year = 2024
var day = 5

func main() {
	input, _ := helpers.GetInput(year, day)

	correctlyOrderedSum := PartOneSolution(input)

	fmt.Printf("Multiplication result: %d.\n", correctlyOrderedSum)
}

func PartOneSolution(input []byte) int {
	var total int

	inputParts := strings.Split(string(input), "\n\n")
	rules, orders := inputParts[0], inputParts[1]

	graph := make(map[int][]int)

	for _, rule := range strings.Split(rules, "\n") {
		parts := strings.Split(rule, "|")
		from, _ := strconv.Atoi(parts[0])
		to, _ := strconv.Atoi(parts[1])

		graph[from] = append(graph[from], to)
	}

	for _, order := range strings.Split(orders, "\n") {
		seen := make(map[int]bool)
		valid := true
		values := strings.Split(order, ",")

		for _, item := range values {
			value, _ := strconv.Atoi(item)

			for _, previous := range graph[value] {
				if seen[previous] {
					valid = false
					break
				}
			}

			seen[value] = true
		}

		if valid {
			midValue, _ := strconv.Atoi(values[len(values)/2])
			total += midValue
		}
	}

	return total
}
