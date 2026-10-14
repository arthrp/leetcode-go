package leetcode

// 85 - Maximal Rectangle
// maximalRectangle returns the area of the largest rectangle of '1' cells in
// matrix.
// Time is O(rows*cols). Extra memory is O(cols).
func maximalRectangle(matrix [][]byte) int {
	if len(matrix) == 0 {
		return 0
	}
	heights := make([]int, len(matrix[0])) // '1's stacked up to this row, per column
	best := 0
	for _, row := range matrix {
		for c, cell := range row {
			if cell == '1' {
				heights[c]++
			} else {
				heights[c] = 0
			}
		}
		// A rectangle whose bottom edge is this row fits under these bars.
		best = max(best, largestRectangleArea(heights))
	}
	return best
}

