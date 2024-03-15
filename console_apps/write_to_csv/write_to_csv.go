package write_to_csv

import (
	"encoding/csv"
	"log"
	"os"
)

func writeToCSV(filename string, data [][]string) {
	file, err := os.Create(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	writer := csv.NewWriter(file)
	defer writer.Flush()
	writer.WriteAll(data)
}
