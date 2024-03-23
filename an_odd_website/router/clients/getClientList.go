package clientRouter

import (
	"encoding/csv"
	"log"
	"os"
	"strconv"
)

func getClientsList() []Client {
	f, e := os.Open("data/clients.csv")
	if e != nil {
		log.Fatal(e)
	}
	defer f.Close()

	r := csv.NewReader(f)
	d, e := r.ReadAll()
	if e != nil {
		log.Fatal(e)
	}

	cl := []Client{}
	for id, item := range d {
		v, e := strconv.ParseBool(item[3])
		if e != nil {
			log.Fatal(e)
		}

		cl = append(cl, Client{
			ClientID:     id,
			FirstName:    item[0],
			LastName:     item[1],
			EmailAddress: item[2],
			Visible:      v,
		})
	}
	return cl
}
