package main

import (
	"adventofcode/internal/helpers"
	"bytes"
	"fmt"
	"os"
	"strconv"
)

var day = 2
var year = 2015

func main() {
	input, err := helpers.GetInput(day, year)
	if err != nil {
		fmt.Printf("Could not get input %s", err)
		os.Exit(1)
	}

	neededArea := Solution(input)
	fmt.Printf("Total area needed: %d", neededArea)
}

func Solution(input []byte) int {
	var neededArea int
	for _, line := range bytes.Split(input, []byte("\n")) {
		dimensions := bytes.Split(line, []byte("x"))
		x, _ := strconv.Atoi(string(dimensions[0]))
		y, _ := strconv.Atoi(string(dimensions[1]))
		z, _ := strconv.Atoi(string(dimensions[2]))
		neededArea += 2*(x*y+y*z+z*x) + min(x*y, y*z, z*x)
	}

	return neededArea
}
