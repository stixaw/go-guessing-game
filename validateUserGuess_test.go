package main

import (
	"testing"
)

func TestValidateUserGuess(t *testing.T) {
	cases := []struct {
		guess    int
		min      int
		max      int
		expected bool
	}{
		{1, 1, 3, true},
		{-1, 1, 3, false},
		{5, 1, 3, false},
	}

	for _, c := range cases {
		output := validateUserGuess(c.guess, c.min, c.max)
		if c.expected != output {
			t.Errorf("Failed ! got %v want %v", output, c.expected)
		} else {
			t.Logf("Success !")
		}
	}
}
