package employeeRouter

import (
	"an_odd_website/router"
	"encoding/json"
	"log"
	"os"
)

func getEmployeeList() []Employee {
	f, e := os.ReadFile("data/employees.json")
	if e != nil {
		log.Fatal(e)
	}

	var el []Employee
	e = json.Unmarshal(f, &el)
	router.CheckAndLogError(e)
	return el
}
