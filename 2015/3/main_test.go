package main

import "testing"

func TestPartOneSolution(t *testing.T) {
	var tests = []struct {
		input []byte
		want  int
	}{
		{[]byte("^"), 2},
		{[]byte("^>v<"), 4},
		{[]byte("^v^v^v^v^v"), 2},
	}
	for _, test := range tests {
		if got := PartOneSolution(test.input); got != test.want {
			t.Errorf("PartOneSolution(%s) = %d, want %d", test.input, got, test.want)
		}
	}
}

func TestPartTwoSolution(t *testing.T) {
	var tests = []struct {
		input []byte
		want  int
	}{
		{[]byte("^v"), 3},
		{[]byte("^>v<"), 3},
		{[]byte("^v^v^v^v^v"), 11},
	}
	for _, test := range tests {
		if got := PartTwoSolution(test.input); got != test.want {
			t.Errorf("PartTwoSolution(%s) = %d, want %d", test.input, got, test.want)
		}
	}
}
