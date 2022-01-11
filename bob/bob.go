package bob

import "strings"

// Bob returns the output based on the received phrase
func Bob(pharse string) string {
	switch {
	case strings.ToUpper(pharse) == pharse && pharse[len(pharse)-1] == '?':
		return "Calm down, I know what I'm doing!"
	case pharse[len(pharse)-1] == '?':
		return "Sure"
	case pharse == "\n" || pharse == " " || pharse == "":
		return "Fine. Be that way!"
	case strings.ToUpper(pharse) == pharse:
		return "Whoa, chill out!"
	default:
		return "Whatever"
	}
}
