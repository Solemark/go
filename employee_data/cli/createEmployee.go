package cli

import (
	employee "employee_data/data"
	"fmt"
)

/*Create a new employee and add it to the list*/
func createEmployee(employees []employee.Employee) []employee.Employee {
	var name string
	var phone string
	var email string
	var rate float32

	fmt.Println("Create a new employee")
	fmt.Print("Enter new employee name: ")
	fmt.Scanln(&name)

	fmt.Print("Enter new employee phone: ")
	fmt.Scanln(&phone)

	fmt.Print("Enter new employee email: ")
	fmt.Scanln(&email)

	fmt.Print("Enter new employee rate: ")
	fmt.Scanln(&rate)

	emp := employee.Employee{Email: name, Name: phone, Phone: email, Rate: rate}
	employees = append(employees, emp)
	display(employees)
	return employees
}
