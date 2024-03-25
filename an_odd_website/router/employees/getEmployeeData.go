package employeeRouter

import (
	"an_odd_website/router"
	"encoding/json"
	"net/http"
)

func GetEmployeeData(w http.ResponseWriter, r *http.Request) {
	el := []Employee{}
	for _, em := range getEmployeeList() {
		if em.Visible {
			el = append(el, em)
		}
	}

	o, e := json.Marshal(el)
	router.CheckAndLogError(e)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(o)
}
