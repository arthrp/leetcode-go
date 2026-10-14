package leetcode

// 188 - Best Time to Buy and Sell Stock IV
// maxProfitIV returns the most profit from at most k buy-then-sell
// transactions that do not overlap.
// Time is O(k*len(prices)). Extra memory is O(k).
func maxProfitIV(k int, prices []int) int {
	// buy[t] holds the t-th share, bought out of sell[t-1]'s profit, and
	// sell[t] has sold it. Sales start at 0, as if that transaction was
	// skipped, so sell[k] also covers fewer than k transactions. prices is
	// never empty, so day 0 can seed every buy.
	buy := make([]int, k+1)
	sell := make([]int, k+1)
	for t := 1; t <= k; t++ {
		buy[t] = -prices[0]
	}
	for _, p := range prices[1:] {
		for t := 1; t <= k; t++ {
			// Each step keeps its best so far or acts today on top of the
			// step before it. In order of t, a step can build on one that
			// already acted today. That trades twice at the same price,
			// which nets zero.
			buy[t] = max(buy[t], sell[t-1]-p)
			sell[t] = max(sell[t], buy[t]+p)
		}
	}
	return sell[k]
}
