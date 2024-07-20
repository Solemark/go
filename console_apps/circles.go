package console_apps

import "math"

func getArea[R Radius](radius R) R {
	switch {
	case radius > 0:
		return math.Pi * (radius * radius)
	default:
		return 0
	}
}

func getCirc[R Radius](radius R) R {
	switch {
	case radius > 0:
		return 2 * math.Pi * radius
	default:
		return 0
	}
}
