package console_apps

import (
	"sort"
)

func numericalSort(data []int) []int {
	sort.Ints(data)
	return data
}
