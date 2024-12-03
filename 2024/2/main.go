package main

import (
	"adventofcode/internal/helpers"
	"bytes"
	"fmt"
	"math"
	"strconv"
)

var year = 2024
var day = 2

func main() {
	input, _ := helpers.GetInput(year, day)

	safeReports := PartOneSolution(input)

	fmt.Printf("Safe reports: %d\n", safeReports)
}

func parseLine(line []byte) []int {
	var values = bytes.Split(line, []byte(" "))
	result := make([]int, len(values))

	for i, val := range values {
		result[i], _ = strconv.Atoi(string(val))
	}

	return result
}

func PartOneSolution(input []byte) int {
	lines := bytes.Split(input, []byte("\n"))
	safeReports := 0

	for _, line := range lines {
		ascending := true
		descending := true

		report := parseLine(line)

		for i, val := range report[1:] {
			if val == report[i] || math.Abs(float64(val-report[i])) > 3 {
				ascending = false
				descending = false
				break
			}

			if val < report[i] {
				ascending = false
			} else {
				descending = false
			}
		}

		if ascending || descending {
			fmt.Println(report)
			safeReports += 1
		}
	}

	return safeReports
}

func PartTwoSolution(input []byte) int {
	lines := bytes.Split(input, []byte("\n"))
	safeReports := 0

	for _, line := range lines {
		ascending := true
		descending := true

		report := parseLine(line)

		for i, val := range report[1:] {
			if val == report[i] || math.Abs(float64(val-report[i])) > 3 {
				ascending = false
				descending = false
				break
			}

			if val < report[i] {
				ascending = false
			} else {
				descending = false
			}
		}

		if ascending || descending {
			fmt.Println(report)
			safeReports += 1
		}
	}

	return safeReports
}
