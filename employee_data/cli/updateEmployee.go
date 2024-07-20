package cli

import (
	employee "employee_data/data"
	"fmt"
)

/*Update an existing employee*/
func updateEmployee(employees []employee.Employee) []employee.Employee {
	name, phone, email, input := "", "", "", ""
	flag := false
	rate := 0.0

	fmt.Print("Enter employee name: ")
	fmt.Scanln(&input)
	for key, employee := range employees {
		if input == employee.Name {
			flag = true
			fmt.Print("Update employee name: ")
			fmt.Scanln(&name)

			fmt.Print("Update employee phone: ")
			fmt.Scanln(&phone)

			fmt.Print("Update employee email: ")
			fmt.Scanln(&email)

			fmt.Print("Update employee rate: ")
			fmt.Scanln(&rate)

			employees[key].Name = name
			employees[key].Phone = phone
			employees[key].Email = email
			employees[key].Rate = rate
		}
	}
	if flag {
		display(employees)
	} else {
		fmt.Println("Error! Couldn't find employee")
	}
	return employees
}
