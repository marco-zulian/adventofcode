package main

import "testing"

func TestPartOneSolution(t *testing.T) {
	var tests = []struct {
		input string
		times int
		want  int
	}{
		{"1", 1, 2},
		{"11", 1, 2},
		{"21", 1, 4},
		{"1211", 1, 6},
		{"111221", 1, 6},
		{"1", 5, 6},
	}

	for _, test := range tests {
		if got := Solution(test.input, test.times); got != test.want {
			t.Errorf("PartOneSolution(%s, %d) = %d, want %d", test.input, test.times, got, test.want)
		}
	}
}
