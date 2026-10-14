package leetcode

import "testing"

func TestGetPermutation(t *testing.T) {
	tests := []struct {
		name string
		n    int
		k    int
		want string
	}{
		{name: "example 1", n: 3, k: 3, want: "213"},
		{name: "example 2", n: 4, k: 9, want: "2314"},
		{name: "example 3", n: 3, k: 1, want: "123"},
		{name: "single digit", n: 1, k: 1, want: "1"},
		{name: "middle permutation", n: 3, k: 4, want: "231"},
		{name: "last permutation", n: 4, k: 24, want: "4321"},
		{name: "first with largest n", n: 9, k: 1, want: "123456789"},
		{name: "last with largest n", n: 9, k: 362880, want: "987654321"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := getPermutation(tt.n, tt.k)
			if got != tt.want {
				t.Fatalf("getPermutation(%d, %d) = %q, want %q", tt.n, tt.k, got, tt.want)
			}
		})
	}
}
