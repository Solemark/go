package console_apps

import "testing"

func TestDatetime(t *testing.T) {
	expect := "the date is Friday the 26th of May 2023"
	answer := checkDate()

	if expect != answer {
		t.Errorf("expected: %s\n recieved: %s", expect, answer)
	}
}
