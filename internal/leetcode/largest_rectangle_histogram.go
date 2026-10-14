package leetcode

// 84 - Largest Rectangle in Histogram
// largestRectangleArea returns the area of the largest rectangle under the
// histogram heights. Each bar is one unit wide.
// Time is O(len(heights)). Extra memory is O(len(heights)).
func largestRectangleArea(heights []int) int {
	n := len(heights)
	stack := make([]int, 0, n) // bar indices with strictly rising heights
	best := 0                  // fits in an int: 1e5 bars times 1e4 height is 1e9
	for i := 0; i <= n; i++ {
		h := 0 // a zero-height bar past the end empties the stack
		if i < n {
			h = heights[i]
		}
		// Bar i is no taller than the top bar, so the top bar's rectangle
		// stops at i-1. If bar i is the same height, its own rectangle covers
		// this one, so cutting this one short loses nothing.
		for len(stack) > 0 && heights[stack[len(stack)-1]] >= h {
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			// Bars between top and i are taller, or they would have popped
			// top. The bar now below top is the nearest shorter bar on its
			// left, so the rectangle starts just after it, or at 0 if there
			// is none.
			left := -1
			if len(stack) > 0 {
				left = stack[len(stack)-1]
			}
			best = max(best, heights[top]*(i-left-1))
		}
		stack = append(stack, i)
	}
	return best
}
