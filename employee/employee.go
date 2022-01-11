package employee

// Employee represents employee entity
type Employee struct {
	Name string
	Age  int
}

// CheckAge check if the Age is more than or less than 22
func CheckAge(name string, age int) (Employee, bool) {
	if age < 22 {
		return Employee{}, false
	}

	return Employee{name, age}, true
}
