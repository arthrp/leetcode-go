package leetcode

import "testing"

func TestMaxProfitIV(t *testing.T) {
	tests := []struct {
		name   string
		k      int
		prices []int
		want   int
	}{
		{name: "example 1", k: 2, prices: []int{2, 4, 1}, want: 2},
		{name: "example 2", k: 2, prices: []int{3, 2, 6, 5, 0, 3}, want: 7},
		{name: "one day", k: 2, prices: []int{5}, want: 0},
		{name: "falling", k: 2, prices: []int{7, 6, 4, 3, 1}, want: 0},
		{name: "one transaction", k: 1, prices: []int{1, 2, 4, 2, 5}, want: 4},
		{name: "every rise", k: 10, prices: []int{1, 2, 4, 2, 5, 7, 2, 4, 9, 0}, want: 15},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := maxProfitIV(tt.k, tt.prices)
			if got != tt.want {
				t.Fatalf("maxProfitIV(%d, %v) = %d, want %d", tt.k, tt.prices, got, tt.want)
			}
		})
	}
}
