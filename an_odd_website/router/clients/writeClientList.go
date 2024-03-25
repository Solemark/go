package clientRouter

import (
	"an_odd_website/router"
	"encoding/json"
	"os"
)

func writeClientList(cl []Client) {
	f, e := os.Create("data/clients.json")
	router.CheckAndLogError(e)
	defer f.Close()

	j, e := json.MarshalIndent(cl, "", "\t")
	router.CheckAndLogError(e)

	f.Write(j)
}
