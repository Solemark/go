package write_to_csv

import (
	"encoding/csv"
	"os"
	"reflect"
	"testing"
)

func constructor() (string, [][]string) {
	return "test.csv", [][]string{
		{"Greeting", "Count to 5", "3rd Header"},
		{"Hello World!", "1, 2, 3, 4, 5", "3rd Col. Data"},
		{"How are you?", "5, 4, 3, 2, 1", "More random data"},
	}
}

func destructor(filename string) {
	os.Remove("test.csv")
}

func TestFileExists(t *testing.T) {
	filename, data := constructor()
	writeToCSV(filename, data)
	_, err := os.Stat(filename)
	if err != nil {
		t.Fatalf("Couldn't find %s file\n%s", filename, err)
	}
	destructor(filename)
}

func TestFileHasData(t *testing.T) {
	filename, expect := constructor()
	writeToCSV(filename, expect)
	input, err := os.Open(filename)
	if err != nil {
		t.Fatal(err)
	}
	reader := csv.NewReader(input)
	result, err := reader.ReadAll()
	if err != nil {
		t.Fatal(err)
	} else {
		if !reflect.DeepEqual(expect, result) {
			t.Log("Expected:\n", expect)
			t.Log("result:\n", result)
			t.Fatalf("Expected array does not match result!")
		}
	}
	destructor(filename)
}
