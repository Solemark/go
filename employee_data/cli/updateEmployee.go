package cli

import (
	employee "employee_data/data"
	"fmt"
)

/*Update an existing employee*/
func updateEmployee(employees []employee.Employee) []employee.Employee {
	var name string
	var phone string
	var email string
	var rate float32
	var input string
	var flag bool

	fmt.Print("Enter employee name: ")
	fmt.Scanln(&input)
	for key, employee := range employees {
		if input == employee.GetName() {
			flag = true
			fmt.Print("Update employee name: ")
			fmt.Scanln(&name)

			fmt.Print("Update employee phone: ")
			fmt.Scanln(&phone)

			fmt.Print("Update employee email: ")
			fmt.Scanln(&email)

			fmt.Print("Update employee rate: ")
			fmt.Scanln(&rate)

			employees[key].SetName(name)
			employees[key].SetPhone(phone)
			employees[key].SetEmail(email)
			employees[key].SetRate(rate)
		}
	}
	if flag {
		display(employees)
	} else {
		fmt.Println("Error! Couldn't find employee")
	}
	return employees
}
