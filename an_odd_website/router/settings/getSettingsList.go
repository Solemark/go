package settingRouter

import (
	"encoding/csv"
	"log"
	"os"
	"strconv"
)

func getSettingsList() []Setting {
	f, e := os.Open("data/settings.csv")
	if e != nil {
		log.Fatal(e)
	}
	defer f.Close()

	r := csv.NewReader(f)
	d, e := r.ReadAll()
	if e != nil {
		log.Fatal(e)
	}

	sl := []Setting{}
	for _, item := range d {
		en, e := strconv.ParseBool(item[1])
		if e != nil {
			log.Fatal(e)
		}

		sl = append(sl, Setting{
			Name:    item[0],
			Enabled: en,
		})
	}
	return sl
}
