package main

import (
	"adventofcode/internal/helpers"
	"bytes"
	"fmt"
	"os"
)

var day = 5
var year = 2015

func main() {
	input, err := helpers.GetInput(day, year)
	if err != nil {
		fmt.Printf("Could not get input %s", err)
		os.Exit(1)
	}

	goodStringsPart1 := PartOneSolution(input)
	goodStringsPart2 := PartTwoSolution(input)
	fmt.Printf("Good strings part 1: %d\nGood strings part 2: %d\n", goodStringsPart1, goodStringsPart2)
}

func PartOneSolution(input []byte) int {
	var goodStrings int

	for _, line := range bytes.Split(input, []byte("\n")) {
		if isGoodStringPart1(string(line)) {
			goodStrings++
		}
	}

	return goodStrings
}

func PartTwoSolution(input []byte) int {
	var goodStrings int

	for _, line := range bytes.Split(input, []byte("\n")) {
		if isGoodStringPart2(string(line)) {
			goodStrings++
		}
	}

	return goodStrings
}

func isGoodStringPart1(input string) bool {
	var vowels int
	var containRepeatedLetters, containForbidenSequence bool
	var forbidenSequences = []string{"ab", "cd", "pq", "xy"}

	runes := []rune(input)
	for i, rn := range runes[:len(runes)-1] {
		if rn == runes[i+1] {
			containRepeatedLetters = true
		}

		if helpers.Contains(string(runes[i:i+2]), forbidenSequences) {
			containForbidenSequence = true
			break
		}

		if helpers.Contains(rn, []rune{'a', 'e', 'i', 'o', 'u'}) {
			vowels++
		}
	}

	if helpers.Contains(runes[len(runes)-1], []rune{'a', 'e', 'i', 'o', 'u'}) {
		vowels++
	}

	return containRepeatedLetters && !containForbidenSequence && vowels >= 3
}

func isGoodStringPart2(input string) bool {
	var pairsStartingIndexes = map[string][]int{}
	var hasRepeatingLetterWithOneRuneBetween, hasNonOverlappingRepeatedPair bool

	runes := []rune(input)
	for i, rn := range runes[:len(runes)-2] {
		pair := string(runes[i : i+2])
		pairsStartingIndexes[pair] = append(pairsStartingIndexes[pair], i)

		if rn == runes[i+2] {
			hasRepeatingLetterWithOneRuneBetween = true
		}
	}

	pair := string(runes[len(runes)-2:])
	pairsStartingIndexes[pair] = append(pairsStartingIndexes[pair], len(runes)-2)

	for _, startingIndexes := range pairsStartingIndexes {
		if len(startingIndexes) > 2 ||
			(len(startingIndexes) == 2 && startingIndexes[1]-startingIndexes[0] > 1) {
			hasNonOverlappingRepeatedPair = true
		}
	}

	return hasNonOverlappingRepeatedPair && hasRepeatingLetterWithOneRuneBetween
}
