package main

import "testing"

func TestLengthOfLongestSubstring(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want int
	}{
		{
			name: "Example 1",
			s:    "abcabcbb",
			want: 3,
		},
		{
			name: "Example 2",
			s:    "bbbbb",
			want: 1,
		},
		{
			name: "Example 3",
			s:    "pwwkew",
			want: 3,
		},
		{
			name: "Empty string",
			s:    "",
			want: 0,
		},
		{
			name: "Single character",
			s:    "a",
			want: 1,
		},
		{
			name: "All unique characters",
			s:    "abcdefg",
			want: 7,
		},
		{
			name: "All duplicates scattered",
			s:    "abba",
			want: 2, // "ab" or "ba"
		},
		{
			name: "Mixed repeating",
			s:    "dvdf",
			want: 3, // "vdf"
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := lengthOfLongestSubstring(tc.s)
			if got != tc.want {
				t.Errorf("FAILED %s: input=%q got=%d want=%d",
					tc.name, tc.s, got, tc.want)
			}
		})
	}
}
