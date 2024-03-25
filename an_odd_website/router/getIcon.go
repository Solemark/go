package router

import (
	"net/http"
	"os"
)

func GetIcon(w http.ResponseWriter, r *http.Request) {
	file, err := os.ReadFile("static/favicon.ico")
	if err != nil {
		ErrorHandler(w, "Error! No Icon!")
		return
	}

	w.Header().Set("Content-Type", "image/x-icon")
	w.WriteHeader(http.StatusOK)
	w.Write(file)
}
