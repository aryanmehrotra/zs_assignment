package main

import "testing"

// TestLeapYear : Tests the leap year funcion
func TestLeapYear(t *testing.T) {
	cases := []struct {
		desc   string // Test Case name
		input  int
		output bool
	}{
		{"divisible by 4", 1992, true},
		{"divisible by 4, 400 and 100", 2000, true},
		{"not divisible by 4 and 400", 1900, false},
	}

	for _, cs := range cases {
		output := checkLeapYear(cs.input)
		// Check if the output is correct or not
		// if not correct return the error using Errorf function of testing package
		if output != cs.output {
			t.Errorf("year %v %v got %v expected %v for leap yaer", cs.input, cs.desc, output, cs.output)
		}
	}
}
