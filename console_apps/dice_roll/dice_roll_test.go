package dice_roll

import "testing"

func getData() []int {
	return []int{
		4, 6, 8, 10, 12, 20,
	}
}

func TestDiceRoll(t *testing.T) {
	var dice []int = getData()

	for _, roll := range dice {
		var answer int = diceRoll(roll)

		if (roll - answer) < 0 {
			t.Error("rolled a: d%i\nrecieved: %i", roll, answer)
		}
	}
}
