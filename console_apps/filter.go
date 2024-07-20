package console_apps

func even(data []int32) []int32 {
	output := []int32{}
	for _, v := range data {
		if v%2 == 0 {
			output = append(output, v)
		}
	}
	return output
}

func odd(data []int32) []int32 {
	output := []int32{}
	for _, v := range data {
		if v%2 != 0 {
			output = append(output, v)
		}
	}
	return output
}

func negative(data []int32) []int32 {
	output := []int32{}
	for _, v := range data {
		if v < 0 {
			output = append(output, v)
		}
	}
	return output
}

func positive(data []int32) []int32 {
	output := []int32{}
	for _, v := range data {
		if v > 0 {
			output = append(output, v)
		}
	}
	return output
}
