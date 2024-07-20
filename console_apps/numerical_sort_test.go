package console_apps

import (
	"reflect"
	"testing"
)

func TestNumericalSort(t *testing.T) {
	data := []int{7, 6, 5, 1, 8, 4, 9, 2, 3, 10}
	expect := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	var result []int = numericalSort(data)

	if !reflect.DeepEqual(expect, result) {
		t.Errorf("Expected: %d\nResult: %d", expect, result)
	}

}
