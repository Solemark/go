package router

import (
	"fmt"
	"net/http"
	"os"
)

func GetScripts(w http.ResponseWriter, r *http.Request) {
	s := r.PathValue("script")
	f, e := os.ReadFile(fmt.Sprintf("static/scripts/%s.js", s))
	errorHandler(w, e)

	w.Header().Set("Content-Type", "text/javascript")
	w.WriteHeader(http.StatusOK)
	w.Write(f)
}
