package console_apps

import (
	"fmt"
	"strings"
)

func Fizzbuzz(fizz int, buzz int, max int) string { return play(fizz, buzz, max, 1, "") }
func play(fizz int, buzz int, max int, i int, output string) string {
	output += ternary((i%fizz == 0), "fizz", "")
	output += ternary((i%buzz == 0), "buzz", "")
	output += ternary((!strings.HasSuffix(output, "z")), fmt.Sprintf("%d", i), "")
	if max <= i {
		return output
	}
	return play(fizz, buzz, max, (i + 1), fmt.Sprintf("%s\n", output))
}
func ternary(condition bool, a string, b string) string {
	if condition {
		return a
	}
	return b
}
