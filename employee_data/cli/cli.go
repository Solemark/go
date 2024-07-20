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
		input := ""
		fmt.Printf("Do you want to:\n%s", getOptions())
		fmt.Scanln(&input)

		switch input {
		case "1":
			employees = createEmployee(employees)
		case "2":
			employees = updateEmployee(employees)
		case "3":
			searchEmployee(employees)
		case "4":
			employees = removeEmployee(employees)
		case "0":
			return
		default:
			fmt.Println("Invalid input!")
		}
	}
}

func getOptions() string {
	return fmt.Sprintln("1. Create\n2. Update\n3. Search\n4. Remove\n0. Exit")
}
