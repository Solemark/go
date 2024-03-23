package settingRouter

import "net/http"

func RemoveSetting(w http.ResponseWriter, r *http.Request) {
	var sl []Setting = getSettingsList()
	var n string = ""

	r.ParseForm()
	for key, value := range r.Form {
		if key == "name" {
			n = value[0]
		}
	}

	for i, s := range sl {
		if s.Name == n {
			sl[i].Enabled = false
		}
	}

	writeSettingList(sl)
	http.Redirect(w, r, "/settings", http.StatusSeeOther)
}
