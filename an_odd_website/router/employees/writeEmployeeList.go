package employeeRouter

import (
	"encoding/csv"
	"log"
	"os"
	"strconv"
)

func writeEmployeeList(data []Employee) {
	f, e := os.Create("data/employees.csv")
	if e != nil {
		log.Fatal(e)
	}
	defer f.Close()

	w := csv.NewWriter(f)
	defer w.Flush()
	el := [][]string{}

	for _, e := range data {
		r := []string{e.FirstName, e.LastName, e.EmailAddress, e.Role, strconv.FormatBool(e.Visible)}
		el = append(el, r)
	}
	w.WriteAll(el)
}
