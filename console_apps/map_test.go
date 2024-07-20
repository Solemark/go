package console_apps

import (
	"reflect"
	"testing"
)

func TestMap(t *testing.T) {
	expect := []int32{3, 6, 9}
	result := Map([]int32{1, 2, 3}, timesThree)
	if !reflect.DeepEqual(expect, result) {
		t.Error("Expected: ", expect, "\nRecieved: ", result)
	}
}
