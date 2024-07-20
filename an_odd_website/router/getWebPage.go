package router

import (
	"fmt"
	"net/http"
	"os"
)

func GetWebPage(w http.ResponseWriter, r *http.Request) {
	url := r.URL.RequestURI()
	var file []byte = []byte{}
	var e error = nil

	file, e = os.ReadFile(fmt.Sprintf("static/%s.html", url))
	if url == "/" {
		file, e = os.ReadFile("static/index.html")
	}
	errorHandler(w, e)

	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusOK)
	w.Write(file)
}
