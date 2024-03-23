package employeeRouter

import (
	"encoding/json"
	"log"
	"os"
)

func writeEmployeeList(el []Employee) {
	f, e := os.Create("data/employees.json")
	if e != nil {
		log.Fatal(e)
	}
	defer f.Close()

	j, e := json.MarshalIndent(el, "", "\t")
	if e != nil {
		log.Fatal(e)
	}

	f.Write(j)
}
