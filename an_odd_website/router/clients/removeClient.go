package clientRouter

import (
	"log"
	"net/http"
	"strconv"
)

func RemoveClient(w http.ResponseWriter, r *http.Request) {
	var cl []Client = getClientList()
	var id string = ""

	r.ParseForm()
	for key, value := range r.Form {
		if key == "client_id" {
			id = value[0]
		}
	}

	i, err := strconv.Atoi(id)
	if err != nil {
		log.Fatal(err)
	}

	cl[i-1].Visible = false
	writeClientList(cl)
	http.Redirect(w, r, "/clients", http.StatusSeeOther)
}
