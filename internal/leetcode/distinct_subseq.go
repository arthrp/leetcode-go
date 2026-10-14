package leetcode

// 115 - Distinct Subsequences
// numDistinct returns how many subsequences of s, told apart by the positions
// they use, spell t.
// Time is O(len(s)*len(t)). Extra memory is O(len(t)).
func numDistinct(s string, t string) int {
	ways := make([]int, len(t)+1) // ways[j] counts subsequences of the s read so far that spell t[:j]
	ways[0] = 1                   // the empty subsequence spells the empty prefix
	for i := 0; i < len(s); i++ {
		// j walks downward so ways[j-1] does not count s[i] yet, which keeps
		// s[i] from filling two positions of t.
		for j := len(t); j > 0; j-- {
			if s[i] == t[j-1] {
				// A spelling of t[:j] either skips s[i], and ways[j] already
				// counts it, or ends with s[i] after a spelling of t[:j-1].
				ways[j] += ways[j-1]
			}
		}
	}
	// Counts for shorter prefixes of t can overflow int, and Go defines signed
	// overflow as wrapping. Every step is an addition, so each count stays
	// right modulo 2^64 (2^32 on 32-bit platforms). The answer fits in a
	// signed 32-bit int, so it comes out exact.
	return ways[len(t)]
}
