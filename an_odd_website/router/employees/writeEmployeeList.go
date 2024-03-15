package employeeRouter

import (
	"encoding/csv"
	"log"
	"os"
	"strconv"
)

func writeEmployeeList(data []Employee) {
	file, err := os.Create("data/employees.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	writer := csv.NewWriter(file)
	defer writer.Flush()
	output := [][]string{}
	for _, c := range data {
		row := []string{c.FirstName, c.LastName, c.EmailAddress, strconv.FormatBool(c.Visible)}
		output = append(output, row)
	}
	writer.WriteAll(output)
}
