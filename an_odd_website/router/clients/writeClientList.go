package clientRouter

import (
	"encoding/json"
	"log"
	"os"
)

func writeClientList(cl []Client) {
	f, e := os.Create("data/clients.json")
	if e != nil {
		log.Fatal(e)
	}
	defer f.Close()

	j, e := json.MarshalIndent(cl, "", "\t")
	if e != nil {
		log.Fatal(e)
	}

	f.Write(j)
}
