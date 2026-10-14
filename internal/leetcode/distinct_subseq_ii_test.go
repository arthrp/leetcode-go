package leetcode

import "testing"

func TestDistinctSubseqII(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want int
	}{
		{name: "all distinct", s: "abc", want: 7},
		{name: "one repeat", s: "aba", want: 6},
		{name: "all same", s: "aaa", want: 3},
		{name: "single letter", s: "a", want: 1},
		{name: "repeated tail", s: "abb", want: 5},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := distinctSubseqII(tt.s)
			if got != tt.want {
				t.Fatalf("distinctSubseqII(%q) = %d, want %d", tt.s, got, tt.want)
			}
		})
	}
}
