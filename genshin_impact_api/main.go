package main

import (
	"encoding/json"
	"fmt"
	"io/fs"
	"log"
	"net/http"
	"os"
	"strings"
)

const jsonContent = "application/json"

func main() {
	http.HandleFunc("/", getBaseRoutes)
	http.HandleFunc("/favicon.ico", getFavIcon)
	http.HandleFunc("/{type}", getTypeRoutes)
	http.HandleFunc("/{type}/{name}", getTypeNamed)

	log.Println("Starting server at localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

func getBaseRoutes(w http.ResponseWriter, r *http.Request) {
	d, e := os.ReadDir("data")
	checkError(w, e)
	o, e := json.Marshal(getNames(d))
	checkError(w, e)
	send_json(w, jsonContent, http.StatusOK, o)
}

func getFavIcon(w http.ResponseWriter, r *http.Request) {
	f, e := os.ReadFile("favicon.ico")
	checkError(w, e)
	w.Header().Set("Content-Type", "image/x-icon")
	w.Write(f)
}

func getTypeRoutes(w http.ResponseWriter, r *http.Request) {
	d, e := os.ReadDir(fmt.Sprintf("data/%s", r.PathValue("type")))
	checkError(w, e)
	o, e := json.Marshal(getNames(d))
	checkError(w, e)
	send_json(w, jsonContent, http.StatusOK, o)
}

func getNames(d []fs.DirEntry) []string {
	o := []string{}
	for _, i := range d {
		o = append(o, strings.ReplaceAll(i.Name(), ".json", ""))
	}
	return o
}

func getTypeNamed(w http.ResponseWriter, r *http.Request) {
	f, e := os.ReadFile(fmt.Sprintf("data/%s/%s.json", r.PathValue("type"), r.PathValue("name")))
	checkError(w, e)
	send_json(w, jsonContent, http.StatusOK, f)
}

func checkError(w http.ResponseWriter, e error) {
	if e != nil {
		s := fmt.Sprintf("{\"error\": \"%s\"}", e.Error())
		log.Println(s)
		send_json(w, jsonContent, http.StatusNotFound, []byte(s))
	}
}

func send_json(w http.ResponseWriter, contentType string, statusCode int, msg []byte) {
	w.Header().Set("Content-Type", contentType)
	if statusCode != 200 {
		w.WriteHeader(statusCode)
	}
	w.Write(msg)
}
