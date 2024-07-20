package console_apps

import (
	"encoding/csv"
	"os"
	"reflect"
	"testing"
)

func TestFileExists(t *testing.T) {
	filename := "test.csv"
	data := [][]string{
		{"Greeting", "Count to 5", "3rd Header"},
		{"Hello World!", "1, 2, 3, 4, 5", "3rd Col. Data"},
		{"How are you?", "5, 4, 3, 2, 1", "More random data"},
	}
	csv := CSV{filename, data}
	csv.write()
	_, err := os.Stat(filename)
	if err != nil {
		t.Fatalf("Couldn't find %s file\n%s", filename, err)
	}
	csv.destroy()
}

func TestFileHasData(t *testing.T) {
	filename := "test.csv"
	data := [][]string{
		{"Greeting", "Count to 5", "3rd Header"},
		{"Hello World!", "1, 2, 3, 4, 5", "3rd Col. Data"},
		{"How are you?", "5, 4, 3, 2, 1", "More random data"},
	}
	c := CSV{filename, data}
	c.write()
	input, err := os.Open(filename)
	if err != nil {
		t.Fatal(err)
	}
	reader := csv.NewReader(input)
	result, err := reader.ReadAll()
	if err != nil {
		t.Fatal(err)
	} else {
		if !reflect.DeepEqual(data, result) {
			t.Log("Expected:\n", data)
			t.Log("result:\n", result)
			t.Fatalf("Expected array does not match result!")
		}
	}
	c.destroy()
}
