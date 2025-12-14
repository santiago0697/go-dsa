package main

import "testing"

func TestIsSubsequence(t *testing.T) {
    tests := []struct {
        name string
        s    string
        t    string
        want bool
    }{
        {
            name: "Example 1",
            s:    "abc",
            t:    "ahbgdc",
            want: true,
        },
        {
            name: "Example 2",
            s:    "axc",
            t:    "ahbgdc",
            want: false,
        },
        {
            name: "Empty s",
            s:    "",
            t:    "abc",
            want: true,
        },
        {
            name: "Empty t",
            s:    "a",
            t:    "",
            want: false,
        },
        {
            name: "Both empty",
            s:    "",
            t:    "",
            want: true,
        },
        {
            name: "Exact match",
            s:    "abc",
            t:    "abc",
            want: true,
        },
        {
            name: "Single character subsequence",
            s:    "a",
            t:    "ba",
            want: true,
        },
        {
            name: "Order matters",
            s:    "aec",
            t:    "abcde",
            want: false,
        },
        {
            name: "Repeated characters in t",
            s:    "aaa",
            t:    "aaabaaa",
            want: true,
        },
        {
            name: "Repeated characters in s not enough in t",
            s:    "aaa",
            t:    "aa",
            want: false,
        },
    }

    for _, tc := range tests {
        t.Run(tc.name, func(t *testing.T) {
            got := isSubsequence(tc.s, tc.t)
            if got != tc.want {
                t.Errorf(
                    "FAILED %s:\n s=%q\n t=%q\n got=%v\n want=%v",
                    tc.name, tc.s, tc.t, got, tc.want,
                )
            }
        })
    }
}
