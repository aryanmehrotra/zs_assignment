package main

import "fmt"

// checkLeapYear : Checks if the year is leap year or not
func checkLeapYear(year int) bool {
	if year%4 == 0 && (year%100 != 0 || year%400 == 0) {
		return true
	}
	return false
}
func main() {
	var year int
	fmt.Println("Enter the year:")
	fmt.Scan(&year)
	print(checkLeapYear(year))
}
