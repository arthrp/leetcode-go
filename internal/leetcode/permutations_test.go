package leetcode

import (
	"fmt"
	"testing"
)

func TestPermute(t *testing.T) {
	tests := []struct {
		nums []int
		want [][]int
	}{
		{
			nums: []int{1, 2, 3},
			want: [][]int{{1, 2, 3}, {1, 3, 2}, {2, 1, 3}, {2, 3, 1}, {3, 1, 2}, {3, 2, 1}},
		},
		{
			nums: []int{0, 1},
			want: [][]int{{0, 1}, {1, 0}},
		},
		{
			nums: []int{1},
			want: [][]int{{1}},
		},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprint(tt.nums), func(t *testing.T) {
			got := permute(append([]int(nil), tt.nums...))
			if !samePermutations(got, tt.want) {
				t.Fatalf("permute(%v) = %v, want %v", tt.nums, got, tt.want)
			}
		})
	}
}

// samePermutations reports whether a and b contain the same rows, ignoring order.
func samePermutations(a, b [][]int) bool {
	if len(a) != len(b) {
		return false
	}
	seen := make(map[string]int, len(a))
	for _, row := range a {
		seen[fmt.Sprint(row)]++
	}
	for _, row := range b {
		key := fmt.Sprint(row)
		if seen[key] == 0 {
			return false
		}
		seen[key]--
	}
	return true
}
