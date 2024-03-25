package employeeRouter

import (
	"an_odd_website/router"
	"net/http"
	"strconv"
)

func UpdateEmployee(w http.ResponseWriter, r *http.Request) {
	var el []Employee = getEmployeeList()
	var id string = ""
	var fn string = ""
	var ln string = ""
	var ea string = ""
	var rl string = ""

	r.ParseForm()
	for key, value := range r.Form {
		if key == "employee_id" {
			id = value[0]
		}
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

	i, e := strconv.Atoi(id)
	router.CheckAndLogError(e)

	el[i-1] = Employee{EmployeeID: i - 1, FirstName: fn, LastName: ln, EmailAddress: ea, Role: rl, Visible: true}
	writeEmployeeList(el)
	http.Redirect(w, r, "/employees", http.StatusSeeOther)
}
