package console_apps

import (
	"strings"
	"testing"
)

func TestGachaRoll(t *testing.T) {
	var data []string = []string{
		"FGO", "AK", "GI",
	}

	for _, value := range data {
		var expect string = GachaRoll(value)
		if !strings.Contains(expect, value) {
			t.Errorf("Expected: %s\nRecieved: %s", expect, value)
		}
	}
}
