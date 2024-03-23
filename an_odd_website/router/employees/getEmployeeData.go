package employeeRouter

import (
	"encoding/json"
	"log"
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
	if e != nil {
		log.Fatal(e)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(o)
}
