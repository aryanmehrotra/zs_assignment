package triangle

import "testing"

// TestTriangle tests the triangle function for different inputs
func TestTriangle(t *testing.T) {
	cases := []struct {
		desc   string // Test case name
		s1     int    // side1 length
		s2     int    // side2 length
		s3     int    // side3 length
		output string
	}{
		{"equilateral", 10, 10, 10, "Equilateral Triangle"},
		{"isosceles", 7, 7, 8, "Isosceles Triangle"},
		{"scalene", 11, 10, 12, "Scalene Triangle"},
		{"not a triangle", 2, 1, 4, "Not a Triangle"},
		{"degenerate", 2, 6, 8, "Degenerate Triangle"},
		{"negative length", -2, 6, 8, "Not a Triangle"},
		{"zero length", 0, 6, 8, "Not a Triangle"},
	}

	for i, tc := range cases {
		output := checkTriangle(tc.s1, tc.s2, tc.s3)

		if output != tc.output {
			t.Errorf("TEST[%d],failed. %s\nExpected : %v \nGot : %v", i, tc.desc, tc.output, output)
		}
	}
}
