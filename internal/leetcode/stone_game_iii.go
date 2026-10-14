package leetcode

// 1406 - Stone Game III
// stoneGameIII reports who wins when Alice and Bob take turns removing 1, 2,
// or 3 stones from the front of stoneValue, Alice first. Each player's score
// is the sum of the values they take. Both play optimally. The result is
// "Alice", "Bob", or "Tie".
//
// The total sum is fixed, so the winner is whoever can force a higher score.
// dp[i] is the best score difference (own score minus the opponent's) that
// the player about to move can force when stoneValue[i:] is still left.
// Taking k stones (k is 1, 2, or 3, and i+k does not run past the end) scores
// sum(stoneValue[i:i+k]) now and leaves the opponent with difference dp[i+k],
// so this move is worth that sum minus dp[i+k]. dp[i] is the maximum of those
// moves. With no stones left, dp[n] is 0.
//
// The scan runs from the end. A move only reads the next three differences,
// so three integers are enough. A positive dp[0] means Alice wins, a negative
// one means Bob wins, and zero is a tie.
//
// Time is O(len(stoneValue)). Extra memory is O(1).
func stoneGameIII(stoneValue []int) string {
	n := len(stoneValue)
	// d1, d2, d3 are dp[i+1], dp[i+2], dp[i+3] as i walks backward.
	// Past the end every difference is 0.
	d1, d2, d3 := 0, 0, 0
	for i := n - 1; i >= 0; i-- {
		// Taking one stone is always legal: at least stoneValue[i] remains.
		best := stoneValue[i] - d1
		if i+1 < n {
			take2 := stoneValue[i] + stoneValue[i+1] - d2
			if take2 > best {
				best = take2
			}
		}
		if i+2 < n {
			take3 := stoneValue[i] + stoneValue[i+1] + stoneValue[i+2] - d3
			if take3 > best {
				best = take3
			}
		}
		// Shift so the next index (i-1) sees dp[i], dp[i+1], dp[i+2].
		d3, d2, d1 = d2, d1, best
	}
	if d1 > 0 {
		return "Alice"
	}
	if d1 < 0 {
		return "Bob"
	}
	return "Tie"
}
