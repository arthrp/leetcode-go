package leetcode

import (
	"strings"
	"testing"
)

func TestNumDistinct(t *testing.T) {
	tests := []struct {
		name string
		s    string
		t    string
		want int
	}{
		{name: "example 1", s: "rabbbit", t: "rabbit", want: 3},
		{name: "example 2", s: "babgbag", t: "bag", want: 5},
		{name: "t longer than s", s: "ab", t: "abc", want: 0},
		{name: "equal strings", s: "abc", t: "abc", want: 1},
		{name: "missing letter", s: "abc", t: "abd", want: 0},
		{name: "case sensitive", s: "A", t: "a", want: 0},
		{name: "choose 2 of 4", s: "aaaa", t: "aa", want: 6},
		{
			// Counts for shorter prefixes of t pass 2^63 (t[:500] alone ends
			// at C(1000, 500)), but the answer is C(1000, 999).
			name: "intermediate counts overflow",
			s:    strings.Repeat("a", 1000),
			t:    strings.Repeat("a", 999),
			want: 1000,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := numDistinct(tt.s, tt.t)
			if got != tt.want {
				t.Fatalf("numDistinct(%q, %q) = %d, want %d", tt.s, tt.t, got, tt.want)
			}
		})
	}
}
