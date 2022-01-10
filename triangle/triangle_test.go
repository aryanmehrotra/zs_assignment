package main

import "testing"

// TestTriangle : tests the triangle function for different inputs
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
	for _, cs := range cases {
		output := checkTriangle(cs.s1, cs.s2, cs.s3)
		// Check if the output is correct or not
		// if not correct return the error using Errorf function of testing package
		if output != cs.output {
			t.Errorf("%v case got %v expected %v", cs.desc, output, cs.output)
		}
	}
}