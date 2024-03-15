package employeeRouter

import (
	"encoding/json"
	"log"
	"net/http"
)

func GetEmployeeData(w http.ResponseWriter, r *http.Request) {
	output := []Employee{}
	for _, client := range getEmployeeList() {
		if client.Visible {
			output = append(output, client)
		}
	}

	out, err := json.Marshal(output)
	if err != nil {
		log.Fatal(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(out)
}
