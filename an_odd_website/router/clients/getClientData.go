package clientRouter

import (
	"an_odd_website/router"
	"encoding/json"
	"net/http"
)

func GetClientData(w http.ResponseWriter, r *http.Request) {
	cl := []Client{}
	for _, c := range getClientList() {
		if c.Visible {
			cl = append(cl, c)
		}
	}

	o, e := json.Marshal(cl)
	router.CheckAndLogError(e)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(o)
}
