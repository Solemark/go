package settingRouter

import (
	"encoding/csv"
	"log"
	"os"
	"strconv"
)

func getSettingsList() []Setting {
	file, err := os.Open("data/settings.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	data, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	output := []Setting{}
	for _, item := range data {
		en, err := strconv.ParseBool(item[1])
		if err != nil {
			log.Fatal(err)
		}

		output = append(output, Setting{
			Name:    item[0],
			Enabled: en,
		})
	}
	return output
}
