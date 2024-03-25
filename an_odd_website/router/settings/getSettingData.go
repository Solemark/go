package settingRouter

import (
	"an_odd_website/router"
	"encoding/json"
	"net/http"
)

func GetSettingData(w http.ResponseWriter, r *http.Request) {
	sl := getSettingList()

	o, e := json.Marshal(sl)
	router.CheckAndLogError(e)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(o)
}
