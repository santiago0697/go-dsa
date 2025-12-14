package main

import "testing"

func TestMajorityElement(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{
			name: "Example 1",
			nums: []int{3, 2, 3},
			want: 3,
		},
		{
			name: "Example 2",
			nums: []int{2, 2, 1, 1, 1, 2, 2},
			want: 2,
		},
		{
			name: "Single element",
			nums: []int{1},
			want: 1,
		},
		{
			name: "All same elements",
			nums: []int{5, 5, 5, 5, 5},
			want: 5,
		},
		{
			name: "Majority at beginning",
			nums: []int{7, 7, 7, 2, 3},
			want: 7,
		},
		{
			name: "Majority at end",
			nums: []int{1, 2, 3, 4, 4, 4, 4},
			want: 4,
		},
		{
			name: "Negative numbers",
			nums: []int{-1, -1, -1, 2, 3},
			want: -1,
		},
		{
			name: "Mixed numbers",
			nums: []int{0, 0, 1, 0, 2, 0, 0},
			want: 0,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := majorityElement(tc.nums)
			if got != tc.want {
				t.Errorf(
					"FAILED %s:\n nums=%v\n got=%d\n want=%d",
					tc.name, tc.nums, got, tc.want,
				)
			}
		})
	}
}
