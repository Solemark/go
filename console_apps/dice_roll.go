package console_apps

import (
	"math/rand"
)

func diceRoll(max int) int {
	return rand.Intn(max) + 1
}
