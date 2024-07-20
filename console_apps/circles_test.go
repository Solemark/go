package console_apps

import (
	"math"
	"testing"
)

func TestCircleArea(t *testing.T) {
	var radii []float64 = []float64{0.0, 1.0, 2.1, -3.2}
	var expect float64
	for _, radius := range radii {
		if radius > 0 {
			expect = math.Pi * (radius * radius)
		} else {
			expect = 0
		}
		answer := getArea(radius)

		if expect != answer {
			t.Errorf("Error! Expected: %g\nrecieved: %g", expect, answer)
		}
	}
}

func TestCircleCirc(t *testing.T) {
	var radii []float64 = []float64{0.0, 1.0, 2.1, -3.2}
	var expect float64
	for _, radius := range radii {
		if radius > 0 {
			expect = 2 * math.Pi * radius
		} else {
			expect = 0
		}
		answer := getCirc(radius)

		if expect != answer {
			t.Errorf("Error! Expected: %g\nrecieved: %g", expect, answer)
		}
	}
}
