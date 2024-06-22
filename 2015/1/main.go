package main

import (
	"adventofcode/internal/helpers"
	"fmt"
)

var year = 2015
var day = 1

func main() {
	var floor int
	var firstBasement int

	input, _ := helpers.GetInput(year, day)

	for i, rn := range input {
		if rn == '(' {
			floor++
		} else {
			floor--
			if floor < 0 && firstBasement == 0 {
				firstBasement = i + 1
			}
		}
	}

	fmt.Printf("Floor at the end: %d.\nFirst step that enters basement: %d.", floor, firstBasement)
}
