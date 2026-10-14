package leetcode

import (
	"fmt"
	"testing"
)

func TestSolveNQueens(t *testing.T) {
	tests := []struct {
		n    int
		want [][]string
	}{
		{
			n:    1,
			want: [][]string{{"Q"}},
		},
		{
			n:    2,
			want: [][]string{},
		},
		{
			n:    3,
			want: [][]string{},
		},
		{
			n: 4,
			want: [][]string{
				{".Q..", "...Q", "Q...", "..Q."},
				{"..Q.", "Q...", "...Q", ".Q.."},
			},
		},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprint(tt.n), func(t *testing.T) {
			got := solveNQueens(tt.n)
			if !sameBoards(got, tt.want) {
				t.Fatalf("solveNQueens(%d) = %v, want %v", tt.n, got, tt.want)
			}
		})
	}
}

func TestSolveNQueensCount(t *testing.T) {
	if got := len(solveNQueens(8)); got != 92 {
		t.Fatalf("len(solveNQueens(8)) = %d, want 92", got)
	}
}

// sameBoards reports whether a and b contain the same boards, ignoring order.
func sameBoards(a, b [][]string) bool {
	if len(a) != len(b) {
		return false
	}
	seen := make(map[string]int, len(a))
	for _, board := range a {
		seen[fmt.Sprint(board)]++
	}
	for _, board := range b {
		key := fmt.Sprint(board)
		if seen[key] == 0 {
			return false
		}
		seen[key]--
	}
	return true
}
