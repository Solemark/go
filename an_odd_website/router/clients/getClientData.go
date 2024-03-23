package clientRouter

import (
	"encoding/json"
	"log"
	"net/http"
)

func GetClientData(w http.ResponseWriter, r *http.Request) {
	cl := []Client{}
	for _, c := range getClientsList() {
		if c.Visible {
			cl = append(cl, c)
		}
	}

	o, e := json.Marshal(cl)
	if e != nil {
		log.Fatal(e)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(o)
}
