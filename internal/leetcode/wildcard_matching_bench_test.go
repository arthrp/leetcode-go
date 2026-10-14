package leetcode

import (
	"strings"
	"testing"
)

// matchKept keeps benchmark results alive so the compiler cannot drop the calls.
var matchKept bool

func BenchmarkIsMatch(b *testing.B) {
	manyStarsS := strings.Repeat("abcde", 200)
	manyStarsP := strings.Repeat("*a*b*c*d*e", 40)
	literal := strings.Repeat("abcdefghij", 80)
	// s is longer than the run of 'a's after '*', and p ends in a 'b' that
	// never occurs, so every placement of the '*' fails. Both algorithms
	// take quadratic time on this input.
	worstS := strings.Repeat("a", 1000)
	worstP := "*" + strings.Repeat("a", 500) + "b"

	cases := []struct {
		name string
		s    string
		p    string
	}{
		{name: "short", s: "adceb", p: "*a*b"},
		{name: "many_stars", s: manyStarsS, p: manyStarsP},
		{name: "literal", s: literal, p: literal},
		{name: "worst", s: worstS, p: worstP},
	}

	for _, tc := range cases {
		b.Run(tc.name, func(b *testing.B) {
			for b.Loop() {
				matchKept = isMatch(tc.s, tc.p)
			}
		})
	}
}
