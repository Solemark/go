package employeeRouter

import (
	"net/http"
)

func NewEmployee(w http.ResponseWriter, r *http.Request) {
	var EmployeeList []Employee = getEmployeeList()
	var firstName string = ""
	var lastName string = ""
	var emailAddress string = ""

	r.ParseForm()
	for key, value := range r.Form {
		if key == "first_name" {
			firstName = value[0]
		}
		if key == "last_name" {
			lastName = value[0]
		}
		if key == "email_address" {
			emailAddress = value[0]
		}
	}

	EmployeeList = append(EmployeeList, Employee{
		EmployeeID:   len(EmployeeList),
		FirstName:    firstName,
		LastName:     lastName,
		EmailAddress: emailAddress,
		Visible:      true,
	})

	writeEmployeeList(EmployeeList)
	http.Redirect(w, r, "/employees", http.StatusSeeOther)
}
