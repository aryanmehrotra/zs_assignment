package twofer

// Twofer check if the given string is a name or "name" literal
func Twofer(s string) string {
	if s == "name" {
		return "One for you, one for me."
	}
	return "One for " + s + " one for me."
}
