package clientRouter

import (
	"encoding/csv"
	"log"
	"os"
	"strconv"
)

func writeClientsList(data []Client) {
	file, err := os.Create("data/clients.csv")
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
