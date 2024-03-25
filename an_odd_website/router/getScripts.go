package router

import (
	"fmt"
	"net/http"
	"os"
)

func GetScripts(w http.ResponseWriter, r *http.Request) {
	script := r.PathValue("script")
	file, err := os.ReadFile(fmt.Sprintf("static/scripts/%s.js", script))
	if err != nil {
		HandleError(w, "Error! Script not found!")
		return
	}

	w.Header().Set("Content-Type", "text/javascript")
	w.WriteHeader(http.StatusOK)
	w.Write(file)
}
