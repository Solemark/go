package calculator

type Number interface {
	int | int8 | int16 | int32 | int64 | float32 | float64
}

func addition[N Number](a N, b N) N       { return a + b }
func subtraction[N Number](a N, b N) N    { return a - b }
func multiplication[N Number](a N, b N) N { return a * b }
func division[N Number](a N, b N) N       { return a / b }
