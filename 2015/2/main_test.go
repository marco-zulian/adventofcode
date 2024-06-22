package main

import "testing"

func TestPartOneSolution(t *testing.T) {
	var tests = []struct {
		input []byte
		want  int
	}{
		{[]byte("2x3x4"), 58},
		{[]byte("1x1x10"), 43},
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
		{[]byte("2x3x4"), 34},
		{[]byte("1x1x10"), 14},
	}
	for _, test := range tests {
		if got := PartTwoSolution(test.input); got != test.want {
			t.Errorf("PartTwoSolution(%s) = %d, want %d", test.input, got, test.want)
		}
	}
}
