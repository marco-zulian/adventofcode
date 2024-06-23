package main

import (
	"testing"
)

func TestPartOneSolution(t *testing.T) {
	var tests = []struct {
		inputKey   string
		inputCount int
		want       int
	}{
		{"abcdef", 5, 609043},
		{"pqrstuv", 5, 1048970},
	}

	for _, test := range tests {
		if got := Solution(test.inputKey, test.inputCount); got != test.want {
			t.Errorf("PartOneSolution(%s, %d) = %d, want %d", test.inputKey, test.inputCount, got, test.want)
		}
	}
}
