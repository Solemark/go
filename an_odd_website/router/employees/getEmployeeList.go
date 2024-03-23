package employeeRouter

import (
	"encoding/csv"
	"log"
	"os"
	"strconv"
)

func getEmployeeList() []Employee {
	f, e := os.Open("data/employees.csv")
	if e != nil {
		log.Fatal(e)
	}

	defer f.Close()
	r := csv.NewReader(f)
	d, e := r.ReadAll()
	if e != nil {
		log.Fatal(e)
	}

	el := []Employee{}
	for id, item := range d {
		vis, err := strconv.ParseBool(item[3])
		if err != nil {
			log.Fatal(err)
		}

		el = append(el, Employee{
			EmployeeID:   id,
			FirstName:    item[0],
			LastName:     item[1],
			EmailAddress: item[2],
			Visible:      vis,
		})
	}
	return el
}
