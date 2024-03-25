package router

import (
	"fmt"
	"net/http"
	"os"
)

func GetWebPage(w http.ResponseWriter, r *http.Request) {
	url := r.URL.RequestURI()
	var file []byte = []byte{}
	var err error = nil

	file, err = os.ReadFile(fmt.Sprintf("static/%s.html", url))
	if url == "/" {
		file, err = os.ReadFile("static/index.html")
	}

	if err != nil {
		ErrorHandler(w, "Error! Page not found!")
		return
	}

	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusOK)
	w.Write(file)
}
