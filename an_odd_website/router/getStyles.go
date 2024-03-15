package router

import (
	"fmt"
	"net/http"
	"os"
)

func GetStyles(w http.ResponseWriter, r *http.Request) {
	style := r.PathValue("style")
	file, err := os.ReadFile(fmt.Sprintf("static/styles/%s.css", style))
	if err != nil {
		handleError(w, "Error! Style not found!")
		return
	}

	w.Header().Set("Content-Type", "text/css")
	w.WriteHeader(http.StatusOK)
	w.Write(file)
}
