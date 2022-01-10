package main

import (
	"fmt"
	"strconv"
)

type result struct {
	ans  string // triangle values for printing
	ansT string // triangle values for testing
}

// printTriangle: Prints the pascal triangle
func printTriangle(n int) {
	output := drawTriangle(n)
	fmt.Println(output.ans)
}

// drawTriangle : Creates the pascal triangle and returns result struct having two value one for testing
// and one for printing the triangle
func drawTriangle(n int) result {
	temp := 1
	ansT := ""
	ans := ""
	// logic to print pascal triangle
	for i := 0; i < n; i++ {
		for j := 1; j <= n-i; j++ {
			ans += " "
		}
		for k := 0; k <= i; k++ {
			if k == 0 || i == 0 {
				temp = 1
			} else {
				temp = temp * (i - k + 1) / k
			}
			ansT = ansT + strconv.Itoa(temp)
			ans = ans + " " + strconv.Itoa(temp)
		}
		ans = ans + "\n"
	}
	answer := result{ans, ansT}
	return answer
}

func main() {
	printTriangle(10)
}
