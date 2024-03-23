package employeeRouter

import (
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
	if e != nil {
		log.Fatal(e)
	}
	return el
}
