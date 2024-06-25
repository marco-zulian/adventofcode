package main

import (
	"adventofcode/internal/helpers"
	"bytes"
	"fmt"
	"os"
	"strconv"
)

var day = 8
var year = 2015

func main() {
	input, err := helpers.GetInput(day, year)
	if err != nil {
		fmt.Printf("Could not get input %s", err)
		os.Exit(1)
	}

	resultPartOne := PartOneSolution(input)
	resultPartTwo := PartTwoSolution(input)
	fmt.Printf("Result Part One: %d\nResult Part Two: %d", resultPartOne, resultPartTwo)
}

func PartOneSolution(input []byte) int {
	var result int
	lines := bytes.Split(input, []byte("\n"))

	for _, line := range lines {
		unquote, _ := strconv.Unquote(string(line))
		result += len(line) - len(unquote)
	}

	return result
}

func PartTwoSolution(input []byte) int {
	var result int
	lines := bytes.Split(input, []byte("\n"))

	for _, line := range lines {
		quoted := strconv.Quote(string(line))
		result += len(quoted) - len(line)
		fmt.Println(quoted, string(line))
	}

	return result
}
