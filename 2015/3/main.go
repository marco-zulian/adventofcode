package main

import (
	"adventofcode/internal/helpers"
	"fmt"
)

func main() {
	input, err := helpers.GetInput(3, 2015)
	if err != nil {
		fmt.Printf("Error getting input: %s", err)
	}

	housesVisited := PartOneSolution(input)
	housesVisitedWithRobot := PartTwoSolution(input)
	fmt.Printf("Total houses visited: %d\n", housesVisited)
	fmt.Printf("Total houses visited with robot: %d\n", housesVisitedWithRobot)
}

func PartOneSolution(input []byte) int {
	var posX, posY int
	visitedPositions := map[string]bool{"0,0": true}

	for _, rn := range input {
		switch rune(rn) {
		case '^':
			posY++
		case 'v':
			posY--
		case '>':
			posX++
		case '<':
			posX--
		}

		visitedPositions[fmt.Sprintf("%d,%d", posX, posY)] = true
	}

	return len(visitedPositions)
}

func PartTwoSolution(input []byte) int {
	var walkingPosX, walkingPosY, waitingPosX, waitingPosY int
	visitedPositions := map[string]bool{"0,0": true}

	for _, rn := range input {
		switch rune(rn) {
		case '^':
			walkingPosY++
		case 'v':
			walkingPosY--
		case '>':
			walkingPosX++
		case '<':
			walkingPosX--
		}

		visitedPositions[fmt.Sprintf("%d,%d", walkingPosX, walkingPosY)] = true
		waitingPosX, walkingPosX = walkingPosX, waitingPosX
		waitingPosY, walkingPosY = walkingPosY, waitingPosY
	}

	return len(visitedPositions)
}
