package employeeRouter

import (
	"log"
	"net/http"
	"strconv"
)

func UpdateEmployee(w http.ResponseWriter, r *http.Request) {
	var el []Employee = getEmployeeList()
	var id string = ""
	var fn string = ""
	var ln string = ""
	var ea string = ""

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
	}

	i, e := strconv.Atoi(id)
	if e != nil {
		log.Fatal(e)
	}

	el[i-1] = Employee{FirstName: fn, LastName: ln, EmailAddress: ea, Visible: true}
	writeEmployeeList(el)
	http.Redirect(w, r, "/employees", http.StatusSeeOther)
}
