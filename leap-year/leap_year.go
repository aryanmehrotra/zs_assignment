package leapyear

// checkLeapYear : Checks if the year is leap year or not
func checkLeapYear(year int) bool {
	if year%4 == 0 && (year%100 != 0 || year%400 == 0) {
		return true
	}
	return false
}
