package console_apps

func palindrome(str string) bool {
	c := len(str) - 1
	for i := 0; i < c; i++ {
		if str[i] != str[c] {
			return false
		}
		c--
	}
	return true
}
