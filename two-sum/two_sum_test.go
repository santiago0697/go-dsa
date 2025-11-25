package main

import "testing"

func TestTwoSum(t *testing.T) {
	tests := []struct {
		input struct {
			nums   []int
			target int
		}
		expected []int
	}{
		{
			input: struct {
				nums   []int
				target int
			}{
				nums:   []int{2, 7, 11, 15},
				target: 9,
			},
			expected: []int{0, 1},
		},
		{
			input: struct {
				nums   []int
				target int
			}{
				nums:   []int{3, 2, 4},
				target: 6,
			},
			expected: []int{1, 2},
		},
		{
			input: struct {
				nums   []int
				target int
			}{
				nums:   []int{3, 3},
				target: 6,
			},
			expected: []int{0, 1},
		},
	}

	for _, test := range tests {
		result := TwoSum(test.input.nums, test.input.target)

		if len(result) != len(test.expected) {
			t.Errorf("Result slice len doesn't match with expected slice len")
		}

		for i, n := range result {
			if n != test.expected[0] && n != test.expected[1] {
				t.Errorf("Result [%d] %d, expected %d", i, n, test.expected[i])
			}
		}
	}
}
