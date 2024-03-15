package dice_roll

import (
	"math/rand"
)

func diceRoll(max int) int {
	return rand.Intn(max) + 1
}
