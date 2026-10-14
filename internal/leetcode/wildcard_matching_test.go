package leetcode

import "testing"

func TestIsMatch(t *testing.T) {
	tests := []struct {
		name string
		s    string
		p    string
		want bool
	}{
		{name: "literal too short", s: "aa", p: "a", want: false},
		{name: "star matches all", s: "aa", p: "*", want: true},
		{name: "question then mismatch", s: "cb", p: "?a", want: false},
		{name: "stars around a literal", s: "adceb", p: "*a*b", want: true},
		{name: "star and question fail", s: "acdcb", p: "a*c?b", want: false},
		{name: "both empty", s: "", p: "", want: true},
		{name: "empty string one star", s: "", p: "*", want: true},
		{name: "empty string many stars", s: "", p: "**", want: true},
		{name: "nonempty string empty pattern", s: "a", p: "", want: false},
		{name: "star in the middle", s: "abc", p: "a*c", want: true},
		{name: "leading star", s: "abc", p: "*c", want: true},
		{name: "backtrack fails", s: "mississippi", p: "m??*ss*?i*pi", want: false},
		{name: "backtrack succeeds", s: "abefcdgiescdfimde", p: "ab*cd?i*de", want: true},
		{name: "leading run of stars", s: "aaaa", p: "***a", want: true},
		{name: "literal before any star", s: "aab", p: "c*a*b", want: false},
		{name: "trailing questions overshoot", s: "abce", p: "abc*??", want: false},
		{name: "empty string one question", s: "", p: "?", want: false},
		{name: "trailing star", s: "a", p: "a*", want: true},
		{name: "questions separated by stars", s: "ab", p: "*?*?*", want: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := isMatch(tt.s, tt.p)
			if got != tt.want {
				t.Fatalf("isMatch(%q, %q) = %v, want %v", tt.s, tt.p, got, tt.want)
			}
		})
	}
}
