package settingRouter

import (
	"an_odd_website/router"
	"encoding/json"
	"log"
	"os"
)

func getSettingList() []Setting {
	f, e := os.ReadFile("data/settings.json")
	router.CheckAndLogError(e)

	var el []Setting
	e = json.Unmarshal(f, &el)
	if e != nil {
		log.Fatal(e)
	}
	return el
}
