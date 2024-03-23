package employeeRouter

import (
	"net/http"
)

func NewEmployee(w http.ResponseWriter, r *http.Request) {
	var el []Employee = getEmployeeList()
	var fn string = ""
	var ln string = ""
	var ea string = ""
	var rl string = ""

	r.ParseForm()
	for key, value := range r.Form {
		if key == "first_name" {
			fn = value[0]
		}
		if key == "last_name" {
			ln = value[0]
		}
		if key == "email_address" {
			ea = value[0]
		}
		if key == "role" {
			rl = value[0]
		}
	}

	el = append(el, Employee{
		EmployeeID:   len(el),
		FirstName:    fn,
		LastName:     ln,
		EmailAddress: ea,
		Role:         rl,
		Visible:      true,
	})

	writeEmployeeList(el)
	http.Redirect(w, r, "/employees", http.StatusSeeOther)
}
