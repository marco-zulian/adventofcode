package main

import "testing"

func TestPartOneSolution(t *testing.T) {
	var tests = []struct {
		input string
		want  int
	}{
		{"ugknbfddgicrmopn", 1},
		{"aaa", 1},
		{"jchzalrnumimnmhp", 0},
		{"haegwjzuvuyypxyu", 0},
		{"dvszwmarrgswjxmb", 0},
		{"ugknbfddgicrmopn\nhaegwjzuvuyypxyu\naaa", 2},
	}

	for _, test := range tests {
		if got := PartOneSolution([]byte(test.input)); got != test.want {
			t.Errorf("PartOneSolution(%s) = %d, want %d", test.input, got, test.want)
		}
	}
}

func TestPartTwoSolution(t *testing.T) {
	var tests = []struct {
		input string
		want  int
	}{
		{"xyxy", 1},
		{"aaa", 0},
		{"abcdefeghi", 0},
		{"abcdefeghiab", 1},
		{"qjhvhtzxzqqjkmpb", 1},
		{"xxyxx", 1},
		{"uurcxstgmygtbstg", 0},
		{"ieodomkazucvgmuy", 0},
		{"xxyxx\nuurcxstgmygtbstg\nxyxy\nabcdefeghiab", 3},
	}

	for _, test := range tests {
		if got := PartTwoSolution([]byte(test.input)); got != test.want {
			t.Errorf("PartTwoSolution(%s) = %d, want %d", test.input, got, test.want)
		}
	}
}
