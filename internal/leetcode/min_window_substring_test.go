package leetcode

import "testing"

func TestMinWindow(t *testing.T) {
	tests := []struct {
		name string
		s    string
		t    string
		want string
	}{
		{name: "example 1", s: "ADOBECODEBANC", t: "ABC", want: "BANC"},
		{name: "example 2", s: "a", t: "a", want: "a"},
		{name: "example 3", s: "a", t: "aa", want: ""},
		{name: "t longer than s", s: "ab", t: "abc", want: ""},
		{name: "case sensitive", s: "a", t: "A", want: ""},
		{name: "window is all of s", s: "abc", t: "cba", want: "abc"},
		{name: "window at the end", s: "bdab", t: "ab", want: "ab"},
		{name: "duplicates in t", s: "aaflslflsldkalskaaa", t: "aaa", want: "aaa"},
		{name: "shrinking past extra characters", s: "cabwefgewcwaefgcf", t: "cae", want: "cwae"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := minWindow(tt.s, tt.t)
			if got != tt.want {
				t.Fatalf("minWindow(%q, %q) = %q, want %q", tt.s, tt.t, got, tt.want)
			}
		})
	}
}
