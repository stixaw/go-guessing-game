package main

import (
	"testing"
)

func TestCompareGuess(t *testing.T) {
	cases := []struct {
		input1   int
		input2   int
		expected bool
	}{
		{1, 5, false},
		{5, 5, true},
	}

	for _, c := range cases {
		output := compareGuess(c.input1, c.input2)

		if c.expected != output {
			t.Errorf("Failed ! got %v want %v", output, c.expected)
		} else {
			t.Logf("Success !")
		}
	}
}
