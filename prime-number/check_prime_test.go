package main

import "testing"

//TestCheckPrime : tests the checkPrime functions implementaion
func TestCheckPrime(t *testing.T) {
	var cases = []struct {
		desc   string // Test case name
		input  int
		output bool
	}{
		{"negative number", -1, false},
		{"zero", 0, false},
		{"one", 1, false},
		{"two", 2, true},
	}
	for i := range cases {
		output := checkPrime(cases[i].input)
		// Check if the output is correct or not
		// if not correct return the error using Errorf function of testing package
		if output != cases[i].output {
			t.Errorf("case with %v got non-prime expected prime", cases[i].desc)
		}
	}
}
