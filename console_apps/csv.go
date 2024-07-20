package console_apps

import (
	"encoding/csv"
	"log"
	"os"
)

type CSV struct {
	filename string
	data     [][]string
}

func (c CSV) write() {
	file, err := os.Create(c.filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	writer := csv.NewWriter(file)
	defer writer.Flush()
	writer.WriteAll(c.data)
}

func (c CSV) destroy() {
	os.Remove(c.filename)
}
