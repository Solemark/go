package console_apps

import "testing"

func Test_fizz_buzz(t *testing.T) {
	const fizz int = 3
	const buzz int = 5
	const max int = 20
	var result string = Fizzbuzz(fizz, buzz, max)
	const expect string = "1\n2\nfizz\n4\nbuzz\nfizz\n7\n8\nfizz\nbuzz\n11\nfizz\n13\n14\nfizzbuzz\n16\n17\nfizz\n19\nbuzz"

	if result != expect {
		t.Errorf("Expected: %s\nRecieved: %s", expect, result)
	}
}
