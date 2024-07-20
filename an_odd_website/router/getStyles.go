package router

import (
	"fmt"
	"net/http"
	"os"
)

func GetStyles(w http.ResponseWriter, r *http.Request) {
	s := r.PathValue("style")
	f, e := os.ReadFile(fmt.Sprintf("static/styles/%s.css", s))
	errorHandler(w, e)

	w.Header().Set("Content-Type", "text/css")
	w.WriteHeader(http.StatusOK)
	w.Write(f)
}
