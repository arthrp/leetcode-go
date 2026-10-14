package leetcode

import "testing"

func TestMaxProfitIII(t *testing.T) {
	tests := []struct {
		name   string
		prices []int
		want   int
	}{
		{name: "example 1", prices: []int{3, 3, 5, 0, 0, 3, 1, 4}, want: 6},
		{name: "example 2", prices: []int{1, 2, 3, 4, 5}, want: 4},
		{name: "example 3", prices: []int{7, 6, 4, 3, 1}, want: 0},
		{name: "one day", prices: []int{5}, want: 0},
		{name: "flat", prices: []int{2, 2, 2}, want: 0},
		{name: "two runs beat one", prices: []int{1, 2, 4, 2, 5, 7, 2, 4, 9, 0}, want: 13},
		{name: "best two of three runs", prices: []int{1, 3, 1, 5, 1, 4}, want: 7},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := maxProfitIII(tt.prices)
			if got != tt.want {
				t.Fatalf("maxProfitIII(%v) = %d, want %d", tt.prices, got, tt.want)
			}
		})
	}
}
