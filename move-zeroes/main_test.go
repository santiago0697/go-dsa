package main

import (
	"reflect"
	"testing"
)

func TestMoveZeroes(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want []int
	}{
		{
			name: "Example 1",
			nums: []int{0, 1, 0, 3, 12},
			want: []int{1, 3, 12, 0, 0},
		},
		{
			name: "All zeros",
			nums: []int{0, 0, 0},
			want: []int{0, 0, 0},
		},
		{
			name: "No zeros",
			nums: []int{1, 2, 3},
			want: []int{1, 2, 3},
		},
		{
			name: "Zeros at the end",
			nums: []int{1, 2, 3, 0, 0},
			want: []int{1, 2, 3, 0, 0},
		},
		{
			name: "Zeros at the start",
			nums: []int{0, 0, 1, 2, 3},
			want: []int{1, 2, 3, 0, 0},
		},
		{
			name: "Alternating zeros",
			nums: []int{0, 1, 0, 2, 0, 3, 0, 4},
			want: []int{1, 2, 3, 4, 0, 0, 0, 0},
		},
		{
			name: "Single element zero",
			nums: []int{0},
			want: []int{0},
		},
		{
			name: "Single element non-zero",
			nums: []int{5},
			want: []int{5},
		},
		{
			name: "Empty array",
			nums: []int{},
			want: []int{},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			numsCopy := make([]int, len(tc.nums))
			copy(numsCopy, tc.nums) // copy so we don't modify original
			moveZeroes(numsCopy)
			if !reflect.DeepEqual(numsCopy, tc.want) {
				t.Errorf("FAILED %s:\n got  %v\n want %v", tc.name, numsCopy, tc.want)
			}
		})
	}
}
