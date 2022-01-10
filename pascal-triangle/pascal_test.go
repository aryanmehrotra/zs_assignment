package main

import (
	"testing"
)

// TestPascal : Test if the values of pascal triangle are correct or not
func TestPascal(t *testing.T) {
	cases := []struct {
		input  int
		output string
	}{
		{0, ""},
		{1, "1"},
		{10, "11112113311464115101051161520156117213535217118285670562881193684126126843691"},
	}

	for _, cs := range cases {
		output := drawTriangle(cs.input)
		// Check if the output is correct or not
		// if not correct return the error using Errorf function of testing package
		if output.ansT != cs.output {
			t.Errorf("expected : %v \n your function returned : %v", cs.output, output.ansT)
		}
	}
}
