package settingRouter

import (
	"encoding/json"
	"log"
	"net/http"
)

func GetSettingData(w http.ResponseWriter, r *http.Request) {
	output := getSettingsList()

	out, err := json.Marshal(output)
	if err != nil {
		log.Fatal(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(out)
}
