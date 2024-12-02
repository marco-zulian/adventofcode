package main

import (
	"adventofcode/internal/helpers"
	"bytes"
	"fmt"
	"math"
	"slices"
	"strconv"
)

var year = 2024
var day = 1

func main() {
	input, _ := helpers.GetInput(year, day)

	distance := PartOneSolution(input)
	similarityScore := PartTwoSolution(input)

	fmt.Printf("Total distance: %d. Similarity: %d\n", distance, similarityScore)
}

func parseInput(input []byte) ([]int, []int) {
	lines := bytes.Split(input, []byte("\n"))
	first := make([]int, len(lines))
	second := make([]int, len(lines))

	for i, line := range lines {
		var values = bytes.Split(line, []byte(" "))
		first[i], _ = strconv.Atoi(string(values[0]))
		second[i], _ = strconv.Atoi(string(values[len(values)-1]))
	}

	slices.Sort(first)
	slices.Sort(second)

	return first, second
}

func PartOneSolution(input []byte) int {
	var distance int
	first, second := parseInput(input)

	for i, v := range first {
		distance += int(math.Abs(float64(v - second[i])))
	}

	return distance
}

func PartTwoSolution(input []byte) int {
	var similarityScore int
	first, second := parseInput(input)
	i, j := 0, 0
	curr := 0

	for i < len(first) && j < len(second) {
		curr = first[i]
		countFirst, countSecond := 0, 0

		for i < len(first) && first[i] == curr {
			countFirst += 1
			i += 1
		}

		for j < len(second) && second[j] <= curr {
			if second[j] == curr {
				countSecond += 1
			}
			j += 1
		}

		similarityScore += countFirst * countSecond * curr
	}

	return similarityScore
}
