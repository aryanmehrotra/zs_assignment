package employee

// Employee : entity which represents employee
type Employee struct {
	name string
	age  int
}

// checkAge check if the age is more than or less than 22
func checkAge(name string, age int) (Employee, bool) {
	if age < 22 {
		return Employee{}, false
	}

	return Employee{name, age}, true
}
