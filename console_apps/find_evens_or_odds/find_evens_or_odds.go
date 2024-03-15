package find_odds_or_evens

func findEvens(data []int) []int {
	output := []int{}
	for _, value := range data {
		if value%2 == 0 {
			output = append(output, value)
		}
	}
	return output
}

func findOdds(data []int) []int {
	output := []int{}
	for _, value := range data {
		if value%2 != 0 {
			output = append(output, value)
		}
	}
	return output
}
