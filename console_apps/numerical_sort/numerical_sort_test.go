package numerical_sort

import (
	"reflect"
	"testing"
)

func getData() ([]int, []int) {
	return []int{
			7, 6, 5, 1, 8, 4, 9, 2, 3, 10,
		}, []int{
			1, 2, 3, 4, 5, 6, 7, 8, 9, 10,
		}
}

func TestNumericalSort(t *testing.T) {
	data, expect := getData()
	var result []int = numericalSort(data)

	if !reflect.DeepEqual(expect, result) {
		t.Errorf("Expected: %d\nResult: %d", expect, result)
	}

}
