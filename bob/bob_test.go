package bob

import "testing"

// TestBob test the bob funtion
func TestBob(t *testing.T) {
	cases := []struct {
		desc   string
		phrase string
		output string
	}{
		{"yells", "AAAAAAAAAAAAAA", "Whoa, chill out!"},
		{"ask question", "How are you?", "Sure"},
		{"yell question", "WHAT ARE YOU DOING?", "Calm down, I know what I'm doing!"},
		{"address without saying anything", " ", "Fine. Be that way!"},
		{"say anything else", " Xy", "Whatever"},
	}

	for i, tc := range cases {
		output := Bob(tc.phrase)

		if output != tc.output {
			t.Errorf("TEST[%d],failed. %s\nExpected : %v \nGot : %v", i, tc.desc, tc.output, output)
		}
	}

}
