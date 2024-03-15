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

func handleError(w http.ResponseWriter, e error) {
	o, e := json.Marshal(fmt.Sprintf("Error! %s", e))
	if e != nil {
		log.Fatal(e)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotFound)
	w.Write(o)
}

func checkError(w http.ResponseWriter, e error) {
	if e != nil {
		handleError(w, e)
	}
}

func getNames(d []fs.DirEntry) []string {
	o := []string{}
	for _, i := range d {
		o = append(o, strings.ReplaceAll(i.Name(), ".json", ""))
	}
	return o
}

func getTypeNamed(w http.ResponseWriter, r *http.Request) {
	t := r.PathValue("type")
	n := r.PathValue("name")
	f, e := os.ReadFile(fmt.Sprintf("data/%s/%s.json", t, n))
	checkError(w, e)

	w.Header().Set("Content-Type", "application/json")
	w.Write(f)
}

func getTypeRoutes(w http.ResponseWriter, r *http.Request) {
	t := r.PathValue("type")
	d, e := os.ReadDir(fmt.Sprintf("data/%s", t))
	checkError(w, e)

	o, e := json.Marshal(getNames(d))
	checkError(w, e)

	w.Header().Set("Content-Type", "application/json")
	w.Write(o)
}

func getBaseRoutes(w http.ResponseWriter, r *http.Request) {
	d, e := os.ReadDir("data")
	checkError(w, e)

	o, e := json.Marshal(getNames(d))
	checkError(w, e)

	w.Header().Set("Content-Type", "application/json")
	w.Write(o)
}

func main() {
	http.HandleFunc("/", getBaseRoutes)
	http.HandleFunc("/{type}", getTypeRoutes)
	http.HandleFunc("/{type}/{name}", getTypeNamed)

	log.Println("Starting server at localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
