package employeeRouter

import (
	"log"
	"net/http"
	"strconv"
)

func RemoveEmployee(w http.ResponseWriter, r *http.Request) {
	var employeeList []Employee = getEmployeeList()
	var id string = ""

	r.ParseForm()
	for key, value := range r.Form {
		if key == "employee_id" {
			id = value[0]
		}
	}

	index, err := strconv.Atoi(id)
	if err != nil {
		log.Fatal(err)
	}

	employeeList[index-1].Visible = false
	writeEmployeeList(employeeList)
	http.Redirect(w, r, "/employees", http.StatusSeeOther)
}
