package main

import (
	"reflect"
	"testing"
)

func buildList(nums []int) *ListNode {
	dummy := &ListNode{}
	curr := dummy
	for _, v := range nums {
		curr.Next = &ListNode{Val: v}
		curr = curr.Next
	}
	return dummy.Next
}

func listToSlice(l *ListNode) []int {
	res := []int{}
	for l != nil {
		res = append(res, l.Val)
		l = l.Next
	}
	return res
}

func TestAddTwoNumbers(t *testing.T) {
	tests := []struct {
		name string
		l1   []int
		l2   []int
		want []int
	}{
		{
			name: "Single digit sum",
			l1:   []int{2},
			l2:   []int{5},
			want: []int{7},
		},
		{
			name: "Overloaded sum",
			l1:   []int{2},
			l2:   []int{9},
			want: []int{1, 1},
		},
		{
			name: "Two digit sum",
			l1:   []int{2, 4},
			l2:   []int{5, 2},
			want: []int{7, 6},
		},
		{
			name: "Two digit sum overloaded",
			l1:   []int{2, 4},
			l2:   []int{5, 9},
			want: []int{7, 3, 1},
		},
		{
			name: "Different length sum",
			l1:   []int{2, 4},
			l2:   []int{5},
			want: []int{7, 4},
		},
		{
			name: "Three digits overloaded on mid",
			l1:   []int{2, 4, 3},
			l2:   []int{5, 6, 4},
			want: []int{7, 0, 8},
		},
		{
			name: "Sum 0 digits",
			l1:   []int{0},
			l2:   []int{0},
			want: []int{0},
		},
		{
			name: "Carry chain",
			l1:   []int{9, 9, 9, 9},
			l2:   []int{1},
			want: []int{0, 0, 0, 0, 1},
		},
		{
			name: "Different lengths",
			l1:   []int{1, 8},
			l2:   []int{0},
			want: []int{1, 8},
		},
		{
			name: "Different lengths",
			l1:   []int{2},
			l2:   []int{1, 8},
			want: []int{3, 8},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {

			l1 := buildList(tc.l1)
			l2 := buildList(tc.l2)
			got := listToSlice(addTwoNumbers(l1, l2))

			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("FAILED %s\n l1=%v l2=%v\n got=%v\nwant=%v",
					tc.name, tc.l1, tc.l2, got, tc.want)
			}
		})
	}
}
