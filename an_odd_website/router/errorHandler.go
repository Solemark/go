package router

import (
	"encoding/json"
	"log"
	"net/http"
)

func handleError(w http.ResponseWriter, message string) {
	o, e := json.Marshal(message)
	CheckAndLogError(e)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotFound)
	w.Write(o)
}

func CheckAndLogError(e error) {
	if e != nil {
		log.Println(e)
	}
}
