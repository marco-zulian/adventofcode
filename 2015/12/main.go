package main

import (
	"adventofcode/internal/helpers"
	"encoding/json"
	"fmt"
	"os"
)

var day = 12
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
	var entries interface{}
	if err := json.Unmarshal(input, &entries); err != nil {
		fmt.Printf("Error unmarshaling JSON %s", err)
		os.Exit(1)
	}

	var result int
	parseJSON(entries, &result)
	return result
}

func parseJSON(input interface{}, counter *int) {
	switch v := input.(type) {
	case []interface{}:
		for _, sv := range v {
			parseJSON(sv, counter)
		}
	case map[string]interface{}:
		for _, sv := range v {
			parseJSON(sv, counter)
		}
	case float64:
		*counter += int(v)
	case string:
		break
	}
}

func PartTwoSolution(input []byte) int {
	var entries interface{}
	if err := json.Unmarshal(input, &entries); err != nil {
		fmt.Printf("Error unmarshaling JSON %s", err)
		os.Exit(1)
	}

	result, _ := parseJSONPartTwo(entries)
	return result
}

func parseJSONPartTwo(input interface{}) (int, bool) {
	var result, sum int
	var ignore bool

	switch v := input.(type) {
	case []interface{}:
		for _, sv := range v {
			sum, _ = parseJSONPartTwo(sv)
			result += sum
		}
	case map[string]interface{}:
		for _, sv := range v {
			if str, ok := sv.(string); ok && str == "red" {
				return 0, true
			}

			sum, ignore = parseJSONPartTwo(sv)
			if ignore {
				return 0, true
			}
			result += sum
		}
	case float64:
		result += int(v)
	case string:
		break
	}

	return result, false
}
