package employeeRouter

import (
	"log"
	"net/http"
	"strconv"
)

func UpdateEmployee(w http.ResponseWriter, r *http.Request) {
	var clientList []Employee = getEmployeeList()
	var id string = ""
	var firstName string = ""
	var lastName string = ""
	var emailAddress string = ""

	r.ParseForm()
	for key, value := range r.Form {
		if key == "employee_id" {
			id = value[0]
		}
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

	index, err := strconv.Atoi(id)
	if err != nil {
		log.Fatal(err)
	}

	clientList[index-1].FirstName = firstName
	clientList[index-1].LastName = lastName
	clientList[index-1].EmailAddress = emailAddress
	writeEmployeeList(clientList)
	http.Redirect(w, r, "/employees", http.StatusSeeOther)
}
