package leapyear

import "testing"

// TestLeapYear : Tests the leap year funcion
func TestLeapYear(t *testing.T) {
	cases := []struct {
		desc   string // Test Case name
		year   int
		okBool bool
	}{
		{"divisible by 4", 1992, true},
		{"divisible by 4, 400 and 100", 2000, true},
		{"not divisible by 4 and 400", 1900, false},
	}

	for i, tc := range cases {
		okBool := checkLeapYear(tc.year)

		// Check if the okBool is correct or not
		// if not correct return the error using Errorf function of testing package
		if okBool != tc.okBool {
			t.Errorf("TEST[%d],failed. %s\nExpected : %v \nGot : %v", i, tc.desc, okBool, tc.okBool)

		}
	}
}
