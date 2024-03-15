package find_odds_or_evens

import (
	"reflect"
	"testing"
)

func getData() []int {
	return []int{
		1, 2, 3, 4, 5, 6, 7, 8, 9, 10,
	}
}

func TestReturnEvens(t *testing.T) {
	var expect []int = []int{2, 4, 6, 8, 10}
	var result []int = findEvens(getData())

	if !reflect.DeepEqual(expect, result) {
		t.Error("Expected: ", expect, "\nRecieved: ", result)
	}
}

func TestReturnOdds(t *testing.T) {
	var expect []int = []int{1, 3, 5, 7, 9}
	var result []int = findOdds(getData())

	if !reflect.DeepEqual(expect, result) {
		t.Error("Expected: ", expect, "\nRecieved: ", result)
	}
}
