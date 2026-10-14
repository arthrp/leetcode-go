package leetcode

import "testing"

func TestLargestRectangleArea(t *testing.T) {
	tests := []struct {
		name    string
		heights []int
		want    int
	}{
		{name: "example 1", heights: []int{2, 1, 5, 6, 2, 3}, want: 10},
		{name: "example 2", heights: []int{2, 4}, want: 4},
		{name: "one bar", heights: []int{5}, want: 5},
		{name: "one zero bar", heights: []int{0}, want: 0},
		{name: "all equal", heights: []int{3, 3, 3}, want: 9},
		{name: "strictly increasing", heights: []int{1, 2, 3, 4, 5}, want: 9},
		{name: "strictly decreasing", heights: []int{5, 4, 3, 2, 1}, want: 9},
		{name: "zero splits bars", heights: []int{2, 0, 2}, want: 2},
		{name: "shortest bar spans all", heights: []int{2, 1, 2}, want: 3},
		{name: "mixed", heights: []int{6, 2, 5, 4, 5, 1, 6}, want: 12},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := largestRectangleArea(tt.heights)
			if got != tt.want {
				t.Fatalf("largestRectangleArea(%v) = %d, want %d", tt.heights, got, tt.want)
			}
		})
	}
}
