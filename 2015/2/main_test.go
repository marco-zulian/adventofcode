package main

import "testing"

func TestMain(t *testing.T) {
	var tests = []struct {
		input []byte
		want  int
	}{
		{[]byte("2x3x4"), 58},
		{[]byte("1x1x10"), 43},
	}
	for _, test := range tests {
		if got := Solution(test.input); got != test.want {
			t.Errorf("Solution(%s) = %d, want %d", test.input, got, test.want)
		}
	}
}
