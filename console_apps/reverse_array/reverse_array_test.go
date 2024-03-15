package reverse_array

import (
	"reflect"
	"testing"
)

func getData() ([]int, []int) {
	return []int{
			1, 2, 3, 4, 5, 6, 7, 8, 9, 10,
		}, []int{
			10, 9, 8, 7, 6, 5, 4, 3, 2, 1,
		}
}

func TestReverseArray(t *testing.T) {
	input, expect := getData()
	result := reverseArray(input)

	if !reflect.DeepEqual(expect, result) {
		t.Errorf("Expect: %d\nResult: %d", result, expect)
		return
	}
}
