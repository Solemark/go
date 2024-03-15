package clientRouter

import (
	"encoding/json"
	"log"
	"net/http"
)

func GetClientData(w http.ResponseWriter, r *http.Request) {
	output := []Client{}
	for _, client := range getClientsList() {
		if client.Visible {
			output = append(output, client)
		}
	}

	out, err := json.Marshal(output)
	if err != nil {
		log.Fatal(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(out)
}
