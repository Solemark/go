package console_apps

import (
	"fmt"
	"time"
)

func checkDate() string {
	datetime := time.Now().Local()
	day := datetime.Weekday().String()
	date := datetime.Day()
	month := datetime.Month().String()
	year := datetime.Year()
	return fmt.Sprintf("the date is %s the %s of %s %d", day, formatString(date), month, year)
}

func formatString(day int) string {
	switch day {
	case 1, 21, 31:
		return fmt.Sprintf("%dst", day)
	case 2, 22:
		return fmt.Sprintf("%dnd", day)
	case 3, 23:
		return fmt.Sprintf("%drd", day)
	default:
		return fmt.Sprintf("%dth", day)
	}
}
