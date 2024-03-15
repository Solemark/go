package cli

import (
	employee "employee_data/data"
	"fmt"
)

/*Find an employee and display them*/
func searchEmployee(employees []employee.Employee) {
	var input string

	fmt.Print("Enter employee name to search: ")
	fmt.Scanln(&input)
	for _, employee := range employees {
		if input == employee.GetName() {
			fmt.Println("Found employee: ", employee)
			return
		}
	}
	fmt.Println("Error! No employee found!")
}
