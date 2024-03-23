package settingRouter

import (
	"encoding/json"
	"log"
	"net/http"
)

func GetSettingData(w http.ResponseWriter, r *http.Request) {
	sl := getSettingsList()

	o, e := json.Marshal(sl)
	if e != nil {
		log.Fatal(e)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(o)
}
