package router

import (
	"encoding/json"
	"log"
	"net/http"
)

func handleError(w http.ResponseWriter, message string) {
	out, err := json.Marshal(message)
	if err != nil {
		log.Fatal(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotFound)
	w.Write(out)
}
