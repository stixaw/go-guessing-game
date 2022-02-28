package main

import (
	"testing"
)

func TestCompareGuess(t *testing.T) {
	input1 := 1
	input2 := 5
	expectedOutput := false

	output := compareGuess(input1, input2)

	if expectedOutput != output {
		t.Errorf("Failed ! got %v want %v", output, expectedOutput)
	} else {
		t.Logf("Success !")
	}

	inputA := 5

	expectedOutput2 := true

	output2 := compareGuess(inputA, input2)

	if expectedOutput2 != output2 {
		t.Errorf("Failed ! got %v want %v", output, expectedOutput)
	} else {
		t.Logf("Success !")
	}
}
