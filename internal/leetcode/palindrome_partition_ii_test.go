package leetcode

import (
	"strings"
	"testing"
)

func TestMinCut(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want int
	}{
		{name: "example 1", s: "aab", want: 1},
		{name: "example 2", s: "a", want: 0},
		{name: "example 3", s: "ab", want: 1},
		{name: "even palindrome", s: "abba", want: 0},
		{name: "all distinct", s: "abcde", want: 4},
		// Taking the longest palindrome prefix first gives aba|a|b, which is
		// 2 cuts. The best split is a|baab.
		{name: "greedy prefix loses", s: "abaab", want: 1},
		{name: "two palindromes", s: "abacdc", want: 1},
		{name: "max length one letter", s: strings.Repeat("a", 2000), want: 0},
		// The first 1999 letters, aba...a, are a palindrome, then b.
		{name: "max length alternating", s: strings.Repeat("ab", 1000), want: 1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := minCut(tt.s)
			if got != tt.want {
				t.Fatalf("minCut(%q) = %d, want %d", tt.s, got, tt.want)
			}
		})
	}
}
