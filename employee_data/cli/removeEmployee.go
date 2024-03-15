package cli

import (
	employee "employee_data/data"
	"fmt"
)

/*Remove the employee from the list*/
func removeEmployee(employees []employee.Employee) []employee.Employee {
	var input string
	var output []employee.Employee

	fmt.Println("Enter name of employee to be deleted: ")
	fmt.Scanln(&input)

	for key, employee := range employees {
		if input != employees[key].GetName() {
			output = append(output, employee)
		}
	}
	display(output)
	return output
}
