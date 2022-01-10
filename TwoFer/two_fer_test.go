package twofer

import "testing"

// TestTwoFer : tests the twoFer function
// takes t from testing package as an input
func TestTwoFer(t *testing.T) {
	cases := []struct {
		desc   string //// Test case name
		input  string
		output string
	}{
		{"with name", "Aryan", "One for Aryan one for me."},
		{"without name", "name", "One for you, one for me."},
	}
	for _, cs := range cases {
		output := twoFer(cs.input)
		// Check if the output is correct or not
		// if not correct return the error using Errorf function of testing package
		if output != cs.output {
			t.Errorf("expected output : %v \n your output : %v", cs.output, output)
		}
	}
}
