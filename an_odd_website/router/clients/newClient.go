package clientRouter

import (
	"net/http"
)

func NewClient(w http.ResponseWriter, r *http.Request) {
	var cl []Client = getClientList()
	var fn string = ""
	var ln string = ""
	var ea string = ""

	r.ParseForm()
	for key, value := range r.Form {
		if key == "first_name" {
			fn = value[0]
		}
		if key == "last_name" {
			ln = value[0]
		}
		if key == "email_address" {
			ea = value[0]
		}
	}

	cl = append(cl, Client{
		ClientID:     len(cl),
		FirstName:    fn,
		LastName:     ln,
		EmailAddress: ea,
		Visible:      true,
	})
	writeClientList(cl)
	http.Redirect(w, r, "/clients", http.StatusSeeOther)
}
