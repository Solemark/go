package settingRouter

import (
	"log"
	"net/http"
	"strconv"
)

func NewSetting(w http.ResponseWriter, r *http.Request) {
	var settingList []Setting = getSettingsList()
	var name string = ""
	var enabled bool = false
	var err error = nil

	r.ParseForm()
	for key, value := range r.Form {
		if key == "name" {
			name = value[0]
		}
		if key == "enabled" {
			enabled, err = strconv.ParseBool(value[0])
			if err != nil {
				log.Fatal(err)
			}
		}
	}

	settingList = append(settingList, Setting{
		Name:    name,
		Enabled: enabled,
	})
	writeSettingList(settingList)
	http.Redirect(w, r, "/settings", http.StatusSeeOther)
}
