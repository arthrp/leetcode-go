package leetcode

import "testing"

func TestMaxPathSum(t *testing.T) {
	tests := []struct {
		name string
		vals []any
		want int
	}{
		{name: "example 1", vals: []any{1, 2, 3}, want: 6},
		{name: "example 2", vals: []any{-10, 9, 20, nil, nil, 15, 7}, want: 42},
		{name: "one negative node", vals: []any{-3}, want: -3},
		{name: "all negative", vals: []any{-2, -1}, want: -1},
		{name: "skip negative branch", vals: []any{1, -2, 3}, want: 4},
		{name: "join through negative root", vals: []any{-1, 5, 6}, want: 10},
		{name: "one-sided chain", vals: []any{5, 4, nil, 3}, want: 12},
		{name: "deep leg and pruned subtree", vals: []any{10, 2, 10, 20, 1, nil, -25, nil, nil, nil, nil, 3, 4}, want: 42},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := maxPathSum(buildTree(tt.vals))
			if got != tt.want {
				t.Fatalf("maxPathSum(%v) = %d, want %d", tt.vals, got, tt.want)
			}
		})
	}
}
