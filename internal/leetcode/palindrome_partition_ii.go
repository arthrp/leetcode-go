package leetcode

// 132 - Palindrome Partitioning II
// minCut returns the fewest cuts that split s into pieces that are all
// palindromes.
// Time is O(n^2) for n = len(s). Extra memory is O(n).
func minCut(s string) int {
	n := len(s)
	// cuts[i] is the fewest cuts that split s[:i] into palindromes. It starts
	// at i-1, a cut between every two letters, since one letter is a
	// palindrome. That makes cuts[0] -1, so a prefix that is a palindrome on
	// its own costs cuts[0]+1 = 0.
	cuts := make([]int, n+1)
	for i := range cuts {
		cuts[i] = i - 1
	}

	// grow visits the palindromes that share a center, shortest first. It
	// starts at s[l..r] and widens it by a letter on each side while the two
	// ends match, so every s[l..r] the loop sees is a palindrome.
	grow := func(l, r int) {
		for ; l >= 0 && r < n && s[l] == s[r]; l, r = l-1, r+1 {
			// The palindrome can be the last piece of s[:r+1], after the best
			// split of s[:l] and one more cut.
			cuts[r+1] = min(cuts[r+1], cuts[l]+1)
		}
	}

	// Every palindrome has one center: a letter if its length is odd, or the
	// gap between two letters if it is even, so growing from every center
	// finds each palindrome once. cuts[l] depends only on palindromes that
	// end before s[l]. Those are centered left of s[c], so the loop has
	// already grown them and cuts[l] is final when grow reads it.
	for c := 0; c < n; c++ {
		grow(c, c)   // odd lengths, centered on s[c]
		grow(c, c+1) // even lengths, centered between s[c] and s[c+1]
	}
	return cuts[n]
}
