package leetcode

import "testing"

func TestTrap(t *testing.T) {
	tests := []struct {
		name   string
		height []int
		want   int
	}{
		{name: "example 1", height: []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}, want: 6},
		{name: "example 2", height: []int{4, 2, 0, 3, 2, 5}, want: 9},
		{name: "one bar", height: []int{1}, want: 0},
		{name: "two equal bars", height: []int{2, 2}, want: 0},
		{name: "strictly increasing", height: []int{1, 2, 3, 4}, want: 0},
		{name: "strictly decreasing", height: []int{4, 3, 2, 1}, want: 0},
		{name: "flat", height: []int{3, 3, 3}, want: 0},
		{name: "single valley", height: []int{5, 0, 5}, want: 5},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := trap(tt.height)
			if got != tt.want {
				t.Fatalf("trap(%v) = %d, want %d", tt.height, got, tt.want)
			}
		})
	}
}
