package settingRouter

import (
	"log"
	"net/http"
	"strconv"
)

func NewSetting(w http.ResponseWriter, r *http.Request) {
	var sl []Setting = getSettingsList()
	var n string = ""
	var en bool = false
	var e error = nil

	r.ParseForm()
	for key, value := range r.Form {
		if key == "name" {
			n = value[0]
		}
		if key == "enabled" {
			en, e = strconv.ParseBool(value[0])
			if e != nil {
				log.Fatal(e)
			}
		}
	}

	sl = append(sl, Setting{
		Name:    n,
		Enabled: en,
	})
	writeSettingList(sl)
	http.Redirect(w, r, "/settings", http.StatusSeeOther)
}
