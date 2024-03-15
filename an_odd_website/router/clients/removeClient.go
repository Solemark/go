package clientRouter

import (
	"log"
	"net/http"
	"strconv"
)

func RemoveClient(w http.ResponseWriter, r *http.Request) {
	var clientList []Client = getClientsList()
	var id string = ""

	r.ParseForm()
	for key, value := range r.Form {
		if key == "client_id" {
			id = value[0]
		}
	}

	index, err := strconv.Atoi(id)
	if err != nil {
		log.Fatal(err)
	}

	clientList[index-1].Visible = false
	writeClientsList(clientList)
	http.Redirect(w, r, "/clients", http.StatusSeeOther)
}
