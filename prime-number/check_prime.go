package primenumber

import (
	"math"
)

// checkPrime checks if the number is prime or not
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
