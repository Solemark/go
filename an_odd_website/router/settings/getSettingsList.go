package settingRouter

import (
	"encoding/json"
	"log"
	"os"
)

func getSettingList() []Setting {
	f, e := os.ReadFile("data/settings.json")
	if e != nil {
		log.Fatal(e)
	}

	var el []Setting
	e = json.Unmarshal(f, &el)
	if e != nil {
		log.Fatal(e)
	}
	return el
}
