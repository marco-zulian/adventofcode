package main

import (
	"fmt"
	"strings"
)

func main() {
	input := "1113122113"
	timesPartOne := 40
	timesPartTwo := 50

	fmt.Println("Result Part 1: ", Solution(input, timesPartOne))
	fmt.Println("Result Part 2: ", Solution(input, timesPartTwo))
}

func Solution(input string, times int) int {
	var builder strings.Builder
	var lastIndex int
	currString := input

	for i := 0; i < times; i++ {
		lastIndex = 0
		for j := 1; j < len(currString); j++ {
			if currString[j] != currString[j-1] {
				builder.WriteString(fmt.Sprintf("%d%s", j-lastIndex, string(currString[j-1])))
				lastIndex = j
			}
		}

		builder.WriteString(fmt.Sprintf("%d%s", len(currString)-lastIndex, string(currString[len(currString)-1])))
		currString = builder.String()
		builder.Reset()
	}

	return len(currString)
}
