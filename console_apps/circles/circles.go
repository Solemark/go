package circles

import "math"

type Number interface {
	float32 | float64
}

func getArea[N Number](radius N) N {
	switch {
	case radius < 0:
		return 0
	case radius > 0:
		return math.Pi * (radius * radius)
	default:
		return 0
	}
}

func getCircumference[N Number](radius N) N {
	switch {
	case radius < 0:
		return 0
	case radius > 0:
		return 2 * math.Pi * radius
	default:
		return 0
	}
}
