package console_apps

import "testing"

func getNumbers() [][]float32 {
	return [][]float32{
		{5.5, 5.5},
		{5.5, -5.5},
		{-5.5, -5.5},
	}
}

func TestAddition(t *testing.T) {
	var numbers [][]float32 = getNumbers()
	var expect float32
	var result float32

	for _, item := range numbers {
		expect = item[0] + item[1]
		result = addition(item[0], item[1])

		if expect != result {
			t.Fatalf("expected: %f\n recieved: %f", expect, result)
		}
	}
}

func TestSubtraction(t *testing.T) {
	var numbers [][]float32 = getNumbers()
	var expect float32
	var result float32

	for _, item := range numbers {
		expect = item[0] - item[1]
		result = subtraction(item[0], item[1])

		if expect != result {
			t.Fatalf("expected: %f\n recieved: %f", expect, result)
		}
	}
}

func TestMultiplication(t *testing.T) {
	var numbers [][]float32 = getNumbers()
	var expect float32
	var result float32

	for _, item := range numbers {
		expect = item[0] * item[1]
		result = multiplication(item[0], item[1])

		if expect != result {
			t.Fatalf("expected: %f\n recieved: %f", expect, result)
		}
	}
}

func TestDivision(t *testing.T) {
	var numbers [][]float32 = getNumbers()
	var expect float32
	var result float32

	for _, item := range numbers {
		expect = item[0] / item[1]
		result = division(item[0], item[1])

		if expect != result {
			t.Fatalf("expected: %f\n recieved: %f", expect, result)
		}
	}
}
