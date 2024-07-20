package cli

import (
	employee "employee_data/data"
	"fmt"
)

/*Create a new employee and add it to the list*/
func createEmployee(employees []employee.Employee) []employee.Employee {
	name, phone, email := "", "", ""
	rate := 0.0

	fmt.Println("Create a new employee")
	fmt.Print("Enter new employee name: ")
	fmt.Scanln(&name)

	fmt.Print("Enter new employee phone: ")
	fmt.Scanln(&phone)

	fmt.Print("Enter new employee email: ")
	fmt.Scanln(&email)

	fmt.Print("Enter new employee rate: ")
	fmt.Scanln(&rate)

	emp := employee.Employee{Name: name, Phone: phone, Email: email, Rate: rate}
	employees = append(employees, emp)
	display(employees)
	return employees
}
