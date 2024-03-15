package clientRouter

import (
	"net/http"
)

func NewClient(w http.ResponseWriter, r *http.Request) {
	var clientList []Client = getClientsList()
	var firstName string = ""
	var lastName string = ""
	var emailAddress string = ""

	r.ParseForm()
	for key, value := range r.Form {
		if key == "first_name" {
			firstName = value[0]
		}
		if key == "last_name" {
			lastName = value[0]
		}
		if key == "email_address" {
			emailAddress = value[0]
		}
	}

	clientList = append(clientList, Client{
		ClientID:     len(clientList),
		FirstName:    firstName,
		LastName:     lastName,
		EmailAddress: emailAddress,
		Visible:      true,
	})
	writeClientsList(clientList)
	http.Redirect(w, r, "/clients", http.StatusSeeOther)
}
