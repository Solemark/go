package map_func

type Number interface {
	int | int8 | int16 | int32 | int64 | float32 | float64
}

func Map[N Number](input []N, mapFunc func(N) N) []N {
	var output []N
	for _, item := range input {
		output = append(output, mapFunc(item))
	}
	return output
}

func timesThree[N Number](x N) N {
	return x * 3
}
