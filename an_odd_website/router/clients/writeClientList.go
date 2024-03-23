package clientRouter

import (
	"encoding/csv"
	"log"
	"os"
	"strconv"
)

func writeClientList(data []Client) {
	f, e := os.Create("data/clients.csv")
	if e != nil {
		log.Fatal(e)
	}
	defer f.Close()

	w := csv.NewWriter(f)
	defer w.Flush()
	cl := [][]string{}

	for _, c := range data {
		r := []string{c.FirstName, c.LastName, c.EmailAddress, strconv.FormatBool(c.Visible)}
		cl = append(cl, r)
	}
	w.WriteAll(cl)
}
