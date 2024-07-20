package router

import (
	"encoding/json"
	"log"
	"net/http"
)

func errorHandler(w http.ResponseWriter, e error) {
	if e != nil {
		o, e := json.Marshal(e)
		CheckAndLogError(e)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		w.Write(o)
	}
}

func CheckAndLogError(e error) {
	if e != nil {
		log.Println(e)
	}
}
