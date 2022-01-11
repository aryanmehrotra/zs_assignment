package twofer

import "testing"

// TestTwoFer : tests the twoFer function
// takes t from testing package as a name
func TestTwoFer(t *testing.T) {
	cases := []struct {
		desc   string //Test case name
		name   string //input
		output string //expected output
	}{
		{"with name", "Aryan", "One for Aryan one for me."},
		{"without name", "name", "One for you, one for me."},
	}

	for i, tc := range cases {
		output := twoFer(tc.name)

		// Check if the output is correct or not
		// if not correct return the error using Errorf function of testing package
		if output != tc.output {
			t.Errorf("TEST[%d],failed. %s\nExpected : %v \nGot : %v", i, tc.desc, tc.output, output)
		}
	}
}
