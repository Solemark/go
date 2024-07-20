package router

import (
	"net/http"
	"os"
)

func GetIcon(w http.ResponseWriter, r *http.Request) {
	f, e := os.ReadFile("static/favicon.ico")
	errorHandler(w, e)

	w.Header().Set("Content-Type", "image/x-icon")
	w.WriteHeader(http.StatusOK)
	w.Write(f)
}
