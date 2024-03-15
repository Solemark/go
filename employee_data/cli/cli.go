package cli

import (
	employee "employee_data/data"
	"fmt"
)

/*Print employee data to console*/
func display(employees []employee.Employee) {
	fmt.Println(employees)
}

/*The runs the CLI*/
func Run() {
	fmt.Println("Employee CLI 1.0")
	employees := []employee.Employee{}
	for {
		input := 0
		fmt.Println("Do you want to 1. CREATE, 2. UPDATE, 3. SEARCH, 4. REMOVE or 0. EXIT")
		fmt.Scanln(&input)

		switch input {
		case 1:
			employees = createEmployee(employees)
		case 2:
			employees = updateEmployee(employees)
		case 3:
			searchEmployee(employees)
		case 4:
			employees = removeEmployee(employees)
		case 0:
			return
		default:
			fmt.Println("Invalid input, please enter 1. CREATE, 2. UPDATE, 3. SEARCH, 4. REMOVE or 0. EXIT")
		}
	}
}
