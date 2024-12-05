package main

import (
	"adventofcode/internal/helpers"
	"fmt"
	"strings"
)

var year = 2024
var day = 4

func main() {
	input, _ := helpers.GetInput(year, day)

	xmasAppearances := PartOneSolution(string(input))
	x_masAppearances := PartTwoSolution(string(input))

	fmt.Printf("XMAS Appearances: %d. X-MAS Appearances: %d\n", xmasAppearances, x_masAppearances)
}

func buildGrid(input string, padding int) [][]string {
	rows := strings.Split(input, "\n")
	grid := make([][]string, len(rows)+2*padding)

	for i := range padding {
		grid[i] = make([]string, len(rows)+2*padding)
	}

	for i, row := range rows {
		grid[i+padding] = make([]string, len(rows)+2*padding)
		for j, cell := range strings.Split(row, "") {
			grid[i+padding][j+padding] = cell
		}
	}

	return grid
}

func PartOneSolution(input string) int {
	PADDING := 3
	var appearances int

	grid := buildGrid(input, PADDING)

	for i := PADDING; i < len(grid)-PADDING; i++ {
		for j := PADDING; j < len(grid[0])-PADDING; j++ {
			if grid[i][j] == "X" {
				if grid[i][j+1] == "M" && grid[i][j+2] == "A" && grid[i][j+3] == "S" {
					appearances++
				}
				if grid[i][j-1] == "M" && grid[i][j-2] == "A" && grid[i][j-3] == "S" {
					appearances++
				}
				if grid[i+1][j] == "M" && grid[i+2][j] == "A" && grid[i+3][j] == "S" {
					appearances++
				}
				if grid[i-1][j] == "M" && grid[i-2][j] == "A" && grid[i-3][j] == "S" {
					appearances++
				}
				if grid[i+1][j+1] == "M" && grid[i+2][j+2] == "A" && grid[i+3][j+3] == "S" {
					appearances++
				}
				if grid[i+1][j-1] == "M" && grid[i+2][j-2] == "A" && grid[i+3][j-3] == "S" {
					appearances++
				}
				if grid[i-1][j+1] == "M" && grid[i-2][j+2] == "A" && grid[i-3][j+3] == "S" {
					appearances++
				}
				if grid[i-1][j-1] == "M" && grid[i-2][j-2] == "A" && grid[i-3][j-3] == "S" {
					appearances++
				}
			}
		}
	}

	return appearances
}

func PartTwoSolution(input string) int {
	PADDING := 1
	var appearances int

	grid := buildGrid(input, PADDING)

	for i := PADDING; i < len(grid)-PADDING; i++ {
		for j := PADDING; j < len(grid[0])-PADDING; j++ {
			if grid[i][j] == "A" {
				if isMAS(grid[i-1][j-1], grid[i][j], grid[i+1][j+1]) && isMAS(grid[i-1][j+1], grid[i][j], grid[i+1][j-1]) {
					appearances++
				}
			}
		}
	}

	return appearances
}

func isMAS(a, b, c string) bool {
	return b == "A" && ((a == "M" && c == "S") || (a == "S" && c == "M"))
}
