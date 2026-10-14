package leetcode

import "testing"

func TestMaximalRectangle(t *testing.T) {
	tests := []struct {
		name   string
		matrix [][]byte
		want   int
	}{
		{
			name: "example 1",
			matrix: [][]byte{
				{'1', '0', '1', '0', '0'},
				{'1', '0', '1', '1', '1'},
				{'1', '1', '1', '1', '1'},
				{'1', '0', '0', '1', '0'},
			},
			want: 6,
		},
		{
			name:   "example 2",
			matrix: [][]byte{{'0'}},
			want:   0,
		},
		{
			name:   "example 3",
			matrix: [][]byte{{'1'}},
			want:   1,
		},
		{
			name: "all ones",
			matrix: [][]byte{
				{'1', '1', '1'},
				{'1', '1', '1'},
			},
			want: 6,
		},
		{
			name: "all zeros",
			matrix: [][]byte{
				{'0', '0', '0'},
				{'0', '0', '0'},
			},
			want: 0,
		},
		{
			name:   "single row",
			matrix: [][]byte{{'1', '1', '0', '1', '1', '1'}},
			want:   3,
		},
		{
			name: "single column",
			matrix: [][]byte{
				{'1'},
				{'1'},
				{'1'},
			},
			want: 3,
		},
		{
			name: "diagonal ones",
			matrix: [][]byte{
				{'1', '0'},
				{'0', '1'},
			},
			want: 1,
		},
		{
			name: "zero breaks a column",
			matrix: [][]byte{
				{'1', '1'},
				{'0', '1'},
				{'1', '1'},
			},
			want: 3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := maximalRectangle(tt.matrix)
			if got != tt.want {
				t.Fatalf("maximalRectangle() = %d, want %d", got, tt.want)
			}
		})
	}
}
