package console_apps

import "testing"

func TestPalindrome(t *testing.T) {
	input := []string{"DAD", "Dad"}
	flag := []bool{true, false}

	for key, str := range input {
		var result bool = palindrome(str)
		if result != flag[key] {
			t.Errorf("\nexpected: %t\nrecieved: %t\nfor: %s = %t", flag[key], result, str, flag[key])
		}
	}
}
