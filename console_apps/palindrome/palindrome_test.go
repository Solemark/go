package palindrome

import "testing"

func getData() ([]string, []bool) {
	return []string{
			"DAD",
			"Dad",
		}, []bool{
			true,
			false,
		}
}

func TestPalindrome(t *testing.T) {
	input, flag := getData()

	for key, str := range input {
		var result bool = palindrome(str)
		if result != flag[key] {
			t.Errorf("\nexpected: %t\nrecieved: %t\nfor: %s = %t", flag[key], result, str, flag[key])
		}
	}
}
