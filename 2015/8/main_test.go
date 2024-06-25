package main

import "testing"

func TestPartOneSolution(t *testing.T) {
	var tests = []struct {
		input []byte
		want  int
	}{
		{[]byte(`""`), 2},
		{[]byte(`"abc"`), 2},
		{[]byte(`"aaa\aaa"`), 3},
		{[]byte(`"\x27"`), 5},
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
		{[]byte(`""`), 4},
		{[]byte(`"abc"`), 4},
		{[]byte(`"aaa\"aaa"`), 6},
		{[]byte(`"\x27"`), 5},
	}
	for _, test := range tests {
		if got := PartTwoSolution(test.input); got != test.want {
			t.Errorf("PartTwoSolution(%s) = %d, want %d", test.input, got, test.want)
		}
	}
}
