package leetcode

import (
	"slices"
	"testing"
)

func TestFindMinII(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{name: "example 1", nums: []int{1, 3, 5}, want: 1},
		{name: "example 2", nums: []int{2, 2, 2, 0, 1}, want: 0},
		{name: "one value", nums: []int{4}, want: 4},
		{name: "all equal", nums: []int{2, 2, 2, 2, 2}, want: 2},
		{name: "tie hides drop after mid", nums: []int{3, 3, 1, 3}, want: 1},
		{name: "tie hides drop before mid", nums: []int{3, 1, 3, 3, 3}, want: 1},
		{name: "one low value among equals", nums: []int{10, 1, 10, 10, 10}, want: 1},
		{name: "min in last slot", nums: []int{2, 3, 4, 1}, want: 1},
		{name: "negative values", nums: []int{-1, 0, 3, -5, -3}, want: -5},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := findMinII(tt.nums)
			if got != tt.want {
				t.Fatalf("findMinII(%v) = %d, want %d", tt.nums, got, tt.want)
			}
		})
	}
}

func TestFindMinIIEveryRotation(t *testing.T) {
	for _, sorted := range [][]int{
		{1, 2, 3, 4, 5, 6, 7, 8},
		{-3, -3, 0, 2, 2, 2, 5, 7, 7},
		{1, 3, 3, 3, 3},
		{1, 1, 1, 1, 1, 2},
	} {
		n := len(sorted)
		for k := 1; k <= n; k++ {
			// Rotating k times moves the last k values to the front.
			nums := slices.Concat(sorted[n-k:], sorted[:n-k])
			if got := findMinII(nums); got != sorted[0] {
				t.Errorf("findMinII(%v) = %d, want %d", nums, got, sorted[0])
			}
		}
	}
}
