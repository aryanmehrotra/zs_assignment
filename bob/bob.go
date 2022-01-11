package bob

import "strings"

// Bob returns the output based on the received phrase
func Bob(phrase string) string {
	switch {
	case strings.ToUpper(phrase) == phrase && phrase[len(phrase)-1] == '?':
		return "Calm down, I know what I'm doing!"
	case phrase[len(phrase)-1] == '?':
		return "Sure"
	case phrase == "\n" || phrase == " " || phrase == "":
		return "Fine. Be that way!"
	case strings.ToUpper(phrase) == phrase:
		return "Whoa, chill out!"
	default:
		return "Whatever"
	}
}
