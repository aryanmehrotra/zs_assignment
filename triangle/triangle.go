package triangle

// checkTriangle : Differentiates and check different triangles based on the length of sides
func checkTriangle(s1, s2, s3 int) string {
	// Check if any side's length is less than 0 or sum of two sides is less than third
	if s1+s2 < s3 || s2+s3 < s1 || s1+s3 < s2 || s1 <= 0 || s2 <= 0 || s3 <= 0 {
		return "Not a Triangle"
		//	Check if sum of two sides are equal to third
	} else if s1+s2 == s3 || s2+s3 == s1 || s1+s3 == s2 {
		return "Degenerate Triangle"
		//	Check if all sides are equal
	} else if s1 == s2 && s2 == s3 {
		return "Equilateral Triangle"
		//	Check if two sides are equal
	} else if s1 == s2 || s2 == s3 || s3 == s1 {
		return "Isosceles Triangle"
	}

	return "Scalene Triangle"
}

func main() {
	print(checkTriangle(10, 11, 12))
}
