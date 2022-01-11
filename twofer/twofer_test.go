package twofer

import "testing"

// TestTwoFer tests the Twofer function
func TestTwoFer(t *testing.T) {
	cases := []struct {
		desc   string
		name   string // input
		output string
	}{
		{"with name", "Aryan", "One for Aryan one for me."},
		{"without name", "name", "One for you, one for me."},
	}

	for i, tc := range cases {
		output := Twofer(tc.name)

		// Check if the output is correct or not
		// if not correct return the error using Errorf function of testing package
		if output != tc.output {
			t.Errorf("TEST[%d],failed. %s\nExpected : %v \nGot : %v", i, tc.desc, tc.output, output)
		}
	}
}
