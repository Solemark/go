package clientRouter

import (
	"encoding/csv"
	"log"
	"os"
	"strconv"
)

func getClientsList() []Client {
	file, err := os.Open("data/clients.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	data, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	output := []Client{}
	for id, item := range data {
		vis, err := strconv.ParseBool(item[3])
		if err != nil {
			log.Fatal(err)
		}

		output = append(output, Client{
			ClientID:     id,
			FirstName:    item[0],
			LastName:     item[1],
			EmailAddress: item[2],
			Visible:      vis,
		})
	}
	return output
}
