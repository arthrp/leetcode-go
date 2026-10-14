package leetcode

// 123 - Best Time to Buy and Sell Stock III
// maxProfitIII returns the most profit from at most two buy-then-sell
// transactions that do not overlap.
// Time is O(len(prices)). Extra memory is O(1).
func maxProfitIII(prices []int) int {
	// Best balance so far after each step: buy1 holds the first share, sell1
	// has sold it, buy2 holds a second share bought out of sell1's profit,
	// and sell2 has sold that too. Sales start at 0, as if the transaction
	// was skipped, so sell2 also covers one transaction or none. prices is
	// never empty, so day 0 can seed both buys.
	buy1, sell1 := -prices[0], 0
	buy2, sell2 := -prices[0], 0
	for _, p := range prices[1:] {
		// Each step keeps its best so far or acts today on top of the step
		// before it. In place, a step can build on one that already acted
		// today. That trades twice at the same price, which nets zero.
		buy1 = max(buy1, -p)
		sell1 = max(sell1, buy1+p)
		buy2 = max(buy2, sell1-p)
		sell2 = max(sell2, buy2+p)
	}
	return sell2
}
