package employee

import "testing"

// TestCheckAge tests checkAge function
func TestCheckAge(t *testing.T) {
	cases := []struct {
		desc   string
		name   string
		age    int
		output Employee
		ok     bool
	}{
		{"less than 22", "Aryan", 20, Employee{}, false},
		{"negative Age", "Aryan", -1, Employee{}, false},
		{"zero Age", "Aryan", 0, Employee{}, false},
		{"more than 22", "Aryan", 23, Employee{"Aryan", 23}, true},
		{"equal to 22", "Aryan", 22, Employee{"Aryan", 22}, true},
	}

	for i, tc := range cases {
		output, ok := checkAge(tc.name, tc.age)

		if output != tc.output {
			t.Errorf("TEST[%d],failed. %s\nExpected : %v \nGot : %v", i, tc.desc, tc.output, output)
		}

		if ok != tc.ok {
			t.Errorf("TEST[%d],failed. %s\nExpected : %v \n Got : %v", i, tc.desc, tc.ok, ok)
		}
	}
}
