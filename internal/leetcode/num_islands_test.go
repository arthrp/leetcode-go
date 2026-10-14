package leetcode

import "testing"

func TestNumIslands(t *testing.T) {
	tests := []struct {
		name string
		grid [][]byte
		want int
	}{
		{
			name: "one island",
			grid: [][]byte{
				{'1', '1', '1', '1', '0'},
				{'1', '1', '0', '1', '0'},
				{'1', '1', '0', '0', '0'},
				{'0', '0', '0', '0', '0'},
			},
			want: 1,
		},
		{
			name: "three islands",
			grid: [][]byte{
				{'1', '1', '0', '0', '0'},
				{'1', '1', '0', '0', '0'},
				{'0', '0', '1', '0', '0'},
				{'0', '0', '0', '1', '1'},
			},
			want: 3,
		},
		{
			name: "diagonals are separate",
			grid: [][]byte{
				{'1', '0'},
				{'0', '1'},
			},
			want: 2,
		},
		{
			name: "all water",
			grid: [][]byte{
				{'0', '0'},
				{'0', '0'},
			},
			want: 0,
		},
		{
			name: "single land",
			grid: [][]byte{{'1'}},
			want: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := numIslands(tt.grid)
			if got != tt.want {
				t.Fatalf("numIslands() = %d, want %d", got, tt.want)
			}
		})
	}
}
