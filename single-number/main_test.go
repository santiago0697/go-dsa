package main

import "testing"

func TestSingleNumber(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{
			name: "Example 1",
			nums: []int{2, 2, 1},
			want: 1,
		},
		{
			name: "Example 2",
			nums: []int{4, 1, 2, 1, 2},
			want: 4,
		},
		{
			name: "Single element",
			nums: []int{7},
			want: 7,
		},
		{
			name: "Negative numbers",
			nums: []int{-1, -1, -2},
			want: -2,
		},
		{
			name: "Mixed positive and negative",
			nums: []int{-5, 3, -5, 4, 4},
			want: 3,
		},
		{
			name: "Zero as single number",
			nums: []int{0, 1, 1},
			want: 0,
		},
		{
			name: "Large numbers",
			nums: []int{100000, 42, 42},
			want: 100000,
		},
		{
			name: "Unordered array",
			nums: []int{9, 1, 9, 2, 2},
			want: 1,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := singleNumber(tc.nums)
			if got != tc.want {
				t.Errorf(
					"FAILED %s:\n nums=%v\n got=%d\n want=%d",
					tc.name, tc.nums, got, tc.want,
				)
			}
		})
	}
}
