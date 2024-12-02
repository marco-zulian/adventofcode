package main

import "testing"

func TestPartOneSolution(t *testing.T) {
	var tests = []struct {
		input []byte
		want  int
	}{
		{[]byte("[1,2,3]"), 6},
		{[]byte(`{"a":2,"b":4}`), 6},
		{[]byte("[[[3]]]"), 3},
		{[]byte(`{"a":{"b":4},"c":-1}`), 3},
		{[]byte(`{"a":[-1,1]}`), 0},
		{[]byte(`[-1,{"a":1}]`), 0},
		{[]byte(`[]`), 0},
		{[]byte(`{}`), 0},
	}

	for _, test := range tests {
		if got := PartOneSolution(test.input); got != test.want {
			t.Errorf("PartOneSolution(%s) = %d, want %d", test.input, got, test.want)
		}
	}
}

func TestPartTwo(t *testing.T) {
	var tests = []struct {
		input []byte
		want  int
	}{
		{[]byte(`[1,2,3]`), 6},
		{[]byte(`[1,{"c":"red","b":2},3]`), 4},
		{[]byte(`{"d":"red","e":[1,2,3,4],"f":5}`), 0},
		{[]byte(`[1,"red",5]`), 6},
		{[]byte(`[1,{"c":"red","b":2},3,"red"]`), 4},
	}

	for _, test := range tests {
		if got := PartTwoSolution(test.input); got != test.want {
			t.Errorf("PartTwoSolution(%s) = %d, want %d", test.input, got, test.want)
		}
	}
}
