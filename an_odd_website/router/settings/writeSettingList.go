package settingRouter

import (
	"encoding/json"
	"log"
	"os"
)

func writeSettingList(sl []Setting) {
	f, e := os.Create("data/settings.json")
	if e != nil {
		log.Fatal(e)
	}
	defer f.Close()

	j, e := json.MarshalIndent(sl, "", "\t")
	if e != nil {
		log.Fatal(e)
	}

	f.Write(j)
}
