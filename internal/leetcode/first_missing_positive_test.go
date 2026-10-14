package leetcode

import "testing"

func TestFirstMissingPositive(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{name: "example 1", nums: []int{1, 2, 0}, want: 3},
		{name: "example 2", nums: []int{3, 4, -1, 1}, want: 2},
		{name: "example 3", nums: []int{7, 8, 9, 11, 12}, want: 1},
		{name: "one is present", nums: []int{1}, want: 2},
		{name: "one is missing", nums: []int{2}, want: 1},
		{name: "duplicate one", nums: []int{1, 1}, want: 2},
		{name: "duplicate two", nums: []int{2, 2}, want: 1},
		{name: "full permutation", nums: []int{3, 1, 2}, want: 4},
		{name: "all non-positive", nums: []int{0, -1, -5}, want: 1},
		{name: "int extremes", nums: []int{-2147483648, 2147483647, 1}, want: 2},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := firstMissingPositive(append([]int(nil), tt.nums...))
			if got != tt.want {
				t.Fatalf("firstMissingPositive(%v) = %d, want %d", tt.nums, got, tt.want)
			}
		})
	}
}
