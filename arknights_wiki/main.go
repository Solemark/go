package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {
	http.HandleFunc("/", list_handler)
	http.HandleFunc("/{operator}", operator_handler)
	log.Println("listening on localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

func list_handler(w http.ResponseWriter, r *http.Request) {
	files, e := os.ReadDir("data")
	str := []string{}
	check_error(w, e)
	for _, v := range files {
		str = append(str, strings.ReplaceAll(v.Name(), ".json", ""))
	}
	o, e := json.Marshal(str)
	check_error(w, e)
	send_json(w, http.StatusOK, o)
}

func operator_handler(w http.ResponseWriter, r *http.Request) {
	o := r.PathValue("operator")
	f, e := os.ReadFile(fmt.Sprintf("data/%s.json", o))
	check_error(w, e)
	send_json(w, http.StatusOK, f)
}

func check_error(w http.ResponseWriter, e error) {
	if e != nil {
		send_json(w, http.StatusForbidden, []byte(fmt.Sprintf("{\"error\":\"%s\"}", e.Error())))
	}
}

func send_json(w http.ResponseWriter, statusCode int, message []byte) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(message)
}
