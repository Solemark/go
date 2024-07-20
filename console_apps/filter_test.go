package console_apps

import (
	"reflect"
	"testing"
)

func TestEvens(t *testing.T) {
	expect := []int32{-4, -2, 0, 2, 4}
	data := []int32{-5, -4, -3, -2, -1, 0, 1, 2, 3, 4, 5}
	result := even(data)

	if !reflect.DeepEqual(expect, result) {
		t.Error("Expected: ", expect, "\nRecieved: ", result)
	}
}

func TestOdds(t *testing.T) {
	expect := []int32{-5, -3, -1, 1, 3, 5}
	data := []int32{-5, -4, -3, -2, -1, 0, 1, 2, 3, 4, 5}
	result := odd(data)

	if !reflect.DeepEqual(expect, result) {
		t.Error("Expected: ", expect, "\nRecieved: ", result)
	}
}

func TestNegative(t *testing.T) {
	expect := []int32{-5, -4, -3, -2, -1}
	data := []int32{-5, -4, -3, -2, -1, 0, 1, 2, 3, 4, 5}
	result := negative(data)

	if !reflect.DeepEqual(expect, result) {
		t.Error("Expected: ", expect, "\nRecieved: ", result)
	}
}

func TestPositive(t *testing.T) {
	expect := []int32{1, 2, 3, 4, 5}
	data := []int32{-5, -4, -3, -2, -1, 0, 1, 2, 3, 4, 5}
	result := positive(data)

	if !reflect.DeepEqual(expect, result) {
		t.Error("Expected: ", expect, "\nRecieved: ", result)
	}
}
