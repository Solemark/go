package gacha_roll

import (
	"fmt"
	"math/rand"
)

func GachaRoll(game string) string {
	switch game {
	case "FGO":
		return roll(100, 300, 5, "FGO", 0, rand.Intn(101))
	case "AK":
		return roll(50, 100, 6, "AK", 0, rand.Intn(51))
	case "GI":
		return roll(60, 90, 5, "GI", 0, rand.Intn(61))
	default:
		return "Error! Unknown game!"
	}
}

func roll(rate int, pity int, rarity int, game string, rolls int, r int) string {
	switch {
	case r == rate:
		return fmt.Sprintf("It took %d rolls to get a %d in %s", rolls, rarity, game)
	case rolls == pity:
		return fmt.Sprintf("You hit pity at %d rolls for your %d in %s", pity, rarity, game)
	}
	return roll(rate, pity, rarity, game, rolls+1, rand.Intn(rate+1))
}
