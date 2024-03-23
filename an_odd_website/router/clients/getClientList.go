package clientRouter

import (
	"encoding/json"
	"log"
	"os"
)

func getClientList() []Client {
	f, e := os.ReadFile("data/clients.json")
	if e != nil {
		log.Fatal(e)
	}

	var cl []Client
	e = json.Unmarshal(f, &cl)
	if e != nil {
		log.Fatal(e)
	}
	return cl
}
