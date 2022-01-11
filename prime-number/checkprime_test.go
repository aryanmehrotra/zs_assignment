package primenumber

import "testing"

//TestCheckPrime tests the CheckPrime functions implementation
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
		{"two", 100, false},
	}

	for i, tc := range cases {
		output := CheckPrime(cases[i].input)

		if output != cases[i].output {
			t.Errorf("TEST[%d],failed. %s\nExpected : %v \nGot : %v", i, tc.desc, tc.output, output)
		}
	}
}
