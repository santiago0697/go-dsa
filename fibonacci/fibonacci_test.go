package main

import "testing"

func TestFibonacci(t *testing.T) {
	tests := []struct {
		input    uint
		expected uint
	}{
		{0, 0},
		{1, 1},
		{5, 5},
		{10, 55},
	}

	for _, test := range tests {
		result := Fibonacci(test.input)

		if result != test.expected {
			t.Errorf("Fibonacci(%d) = %d, expected %d", test.input, result, test.expected)
		}
	}
}
