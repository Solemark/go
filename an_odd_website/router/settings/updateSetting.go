package settingRouter

import (
	"log"
	"net/http"
	"strconv"
)

func UpdateSetting(w http.ResponseWriter, r *http.Request) {
	var sl []Setting = getSettingsList()
	var n string = ""
	var e bool = false
	var er error = nil

	r.ParseForm()
	for key, value := range r.Form {
		if key == "name" {
			n = value[0]
		}
		if key == "enabled" {
			e, er = strconv.ParseBool(value[0])
			if er != nil {
				log.Fatal(er)
			}
		}
	}

	for i, s := range sl {
		if s.Name == n {
			sl[i] = Setting{Name: n, Enabled: e}
		}
	}
	writeSettingList(sl)
	http.Redirect(w, r, "/settings", http.StatusSeeOther)
}
