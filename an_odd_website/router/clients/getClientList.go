package clientRouter

import (
	"an_odd_website/router"
	"encoding/json"
	"os"
)

func getClientList() []Client {
	f, e := os.ReadFile("data/clients.json")
	router.CheckAndLogError(e)

	var cl []Client
	e = json.Unmarshal(f, &cl)
	router.CheckAndLogError(e)
	return cl
}
