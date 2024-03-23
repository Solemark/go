package settingRouter

import (
	"encoding/csv"
	"log"
	"os"
	"strconv"
)

func writeSettingList(data []Setting) {
	file, err := os.Create("data/settings.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()
	output := [][]string{}

	for _, s := range data {
		row := []string{s.Name, strconv.FormatBool(s.Enabled)}
		output = append(output, row)
	}
	writer.WriteAll(output)
}
