package settingRouter

import (
	"encoding/csv"
	"log"
	"os"
	"strconv"
)

func writeSettingList(d []Setting) {
	f, e := os.Create("data/settings.csv")
	if e != nil {
		log.Fatal(e)
	}
	defer f.Close()

	w := csv.NewWriter(f)
	defer w.Flush()
	sl := [][]string{}

	for _, s := range d {
		r := []string{s.Name, strconv.FormatBool(s.Enabled)}
		sl = append(sl, r)
	}
	w.WriteAll(sl)
}
