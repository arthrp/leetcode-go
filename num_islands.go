package leetcode

// 200 - Number of Islands
// numIslands counts groups of land ('1') connected horizontally or vertically.
//
// It scans the grid once. Each unvisited land cell starts a new island. A
// depth-first search then walks every cell of that island and sinks it by
// writing '0', so later scans do not count the same land again. Water and
// cells outside the grid stop the walk. Diagonal neighbors are not connected.
func numIslands(grid [][]byte) int {
	count := 0
	for r := range grid {
		for c := range grid[r] {
			if grid[r][c] != '1' {
				continue
			}
			count++
			sink(grid, r, c)
		}
	}
	return count
}

// sink marks every land cell reachable from (r, c) as water.
func sink(grid [][]byte, r, c int) {
	if r < 0 || c < 0 || r >= len(grid) || c >= len(grid[r]) || grid[r][c] != '1' {
		return
	}
	grid[r][c] = '0'
	sink(grid, r-1, c)
	sink(grid, r+1, c)
	sink(grid, r, c-1)
	sink(grid, r, c+1)
}
