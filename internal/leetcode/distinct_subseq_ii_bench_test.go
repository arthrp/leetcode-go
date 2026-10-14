package leetcode

import (
	"strings"
	"testing"
)

// subseqKept keeps benchmark results alive so the compiler cannot drop the calls.
var subseqKept int

func BenchmarkDistinctSubseqII(b *testing.B) {
	cases := []struct {
		name string
		s    string
	}{
		{name: "short", s: "abcab"},
		{name: "alphabet_2000", s: strings.Repeat("abcdefghijklmnopqrstuvwxyz", 77)[:2000]},
		{name: "same_2000", s: strings.Repeat("a", 2000)},
		{name: "alphabet_100k", s: strings.Repeat("abcdefghijklmnopqrstuvwxyz", 3847)[:100_000]},
	}

	for _, tc := range cases {
		b.Run(tc.name, func(b *testing.B) {
			for b.Loop() {
				subseqKept = distinctSubseqII(tc.s)
			}
		})
	}
}
