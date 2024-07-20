package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", greeting)
	log.Println("listening on localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func greeting(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	u, p, o := r.BasicAuth()
	if !o {
		w.Header().Add("WWW-Authenticate", `Basic realm="Give username and password"`)
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte(`{"message": "No basic auth present"}`))
		return
	}

	if !isAuthorised(u, p) {
		w.Header().Add("WWW-Authenticate", `Basic realm="Give username and password"`)
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte(`{"message": "Invalid username or password"}`))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "welcome to golang world!"}`))
}

func isAuthorised(u string, p string) bool {
	pass, o := getUsers()[u]
	if !o {
		return false
	}
	return p == pass
}

/*
Returns a map of username and password where key = username and value = password
*/
func getUsers() map[string]string {
	return map[string]string{
		"admin": "admin",
	}
}
