package main

import (
	"fmt"
	"math"
)

// checkPrime : checks if the number is prime or not
func checkPrime(input int) bool {
	if input <= 1 {
		return false
	}
	n := int(math.Sqrt(float64(input)))
	for i := 2; i <= n; i++ {
		if input%i == 0 {
			return false
		}
	}
	return true
}
func main() {
	var num int
	fmt.Println("Enter a number:")
	fmt.Scan(&num) // takes input from the user
	fmt.Println(checkPrime(num))
}
