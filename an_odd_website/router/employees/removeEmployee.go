package employeeRouter

import (
	"an_odd_website/router"
	"net/http"
	"strconv"
)

func RemoveEmployee(w http.ResponseWriter, r *http.Request) {
	var el []Employee = getEmployeeList()
	var id string = ""

	r.ParseForm()
	for key, value := range r.Form {
		if key == "employee_id" {
			id = value[0]
		}
	}

	i, e := strconv.Atoi(id)
	router.CheckAndLogError(e)

	el[i-1].Visible = false
	writeEmployeeList(el)
	http.Redirect(w, r, "/employees", http.StatusSeeOther)
}
