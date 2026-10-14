package leetcode

import "testing"

func TestStoneGameIII(t *testing.T) {
	tests := []struct {
		name       string
		stoneValue []int
		want       string
	}{
		{name: "bob wins", stoneValue: []int{1, 2, 3, 7}, want: "Bob"},
		{name: "alice takes three", stoneValue: []int{1, 2, 3, -9}, want: "Alice"},
		{name: "tie on three", stoneValue: []int{1, 2, 3, 6}, want: "Tie"},
		{name: "one stone", stoneValue: []int{1}, want: "Alice"},
		{name: "two stones", stoneValue: []int{1, 2}, want: "Alice"},
		{name: "all negative", stoneValue: []int{-1, -2, -3}, want: "Tie"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := stoneGameIII(tt.stoneValue)
			if got != tt.want {
				t.Fatalf("stoneGameIII(%v) = %q, want %q", tt.stoneValue, got, tt.want)
			}
		})
	}
}
