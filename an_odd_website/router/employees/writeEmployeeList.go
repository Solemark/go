package employeeRouter

import (
	"an_odd_website/router"
	"encoding/json"
	"os"
)

func writeEmployeeList(el []Employee) {
	f, e := os.Create("data/employees.json")
	router.CheckAndLogError(e)
	defer f.Close()

	j, e := json.MarshalIndent(el, "", "\t")
	router.CheckAndLogError(e)

	f.Write(j)
}
