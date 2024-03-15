package numerical_sort

import (
	"sort"
)

func numericalSort(data []int) []int {
	sort.Ints(data)
	return data
}
