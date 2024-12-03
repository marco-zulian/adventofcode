package main

import (
	"adventofcode/internal/helpers"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

var year = 2024
var day = 3

func main() {
	input, _ := helpers.GetInput(year, day)

	total := PartOneSolution(input)
	restrictedTotal := PartTwoSolution(string(input))

	fmt.Printf("Multiplication result: %d. Restricted multiplication result: %d\n", total, restrictedTotal)
}

func PartOneSolution(input []byte) int {
	var total int

	re := regexp.MustCompile(`mul\(\d{1,3},\d{1,3}\)`)

	for _, match := range re.FindAllString(string(input), -1) {
		total += multiply(match)
	}

	return total
}

func PartTwoSolution(input string) int {
	var total int

	doRegex := regexp.MustCompile(`do\(\)`)
	dontRegex := regexp.MustCompile(`don\'t\(\)`)
	multsRegex := regexp.MustCompile(`mul\(\d{1,3},\d{1,3}\)`)

	doIndexes := doRegex.FindAllStringIndex(input, -1)
	dontIndexes := dontRegex.FindAllStringIndex(input, -1)
	multsIndexes := multsRegex.FindAllStringIndex(input, -1)

	var j, k int
	doIndex := -1
	dontIndex := -1

	for i := 0; i < len(multsIndexes); i++ {
		for j < len(doIndexes) {
			if doIndexes[j][0] > multsIndexes[i][0] {
				break
			}
			doIndex = doIndexes[j][0]
			j++
		}

		for k < len(dontIndexes) {
			if dontIndexes[k][0] > multsIndexes[i][0] {
				break
			}
			dontIndex = dontIndexes[k][0]
			k++
		}

		if doIndex >= dontIndex {
			total += multiply(input[multsIndexes[i][0]:multsIndexes[i][1]])
		}
	}

	return total
}

func multiply(input string) int {
	multiplicationResult := 1
	nonNumericRegex := regexp.MustCompile(`[^\d]+`)

	for _, number := range strings.Split(input, ",") {
		integerValue, _ := strconv.Atoi(nonNumericRegex.ReplaceAllString(string(number), ""))

		multiplicationResult *= integerValue
	}

	return multiplicationResult
}
