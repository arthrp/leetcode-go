package leetcode

import (
	"slices"
	"testing"
)

func TestFullJustify(t *testing.T) {
	tests := []struct {
		name     string
		words    []string
		maxWidth int
		want     []string
	}{
		{
			name:     "example 1",
			words:    []string{"This", "is", "an", "example", "of", "text", "justification."},
			maxWidth: 16,
			want: []string{
				"This    is    an",
				"example  of text",
				"justification.  ",
			},
		},
		{
			name:     "example 2",
			words:    []string{"What", "must", "be", "acknowledgment", "shall", "be"},
			maxWidth: 16,
			want: []string{
				"What   must   be",
				"acknowledgment  ",
				"shall be        ",
			},
		},
		{
			name:     "example 3",
			words:    []string{"Science", "is", "what", "we", "understand", "well", "enough", "to", "explain", "to", "a", "computer.", "Art", "is", "everything", "else", "we", "do"},
			maxWidth: 20,
			want: []string{
				"Science  is  what we",
				"understand      well",
				"enough to explain to",
				"a  computer.  Art is",
				"everything  else  we",
				"do                  ",
			},
		},
		{
			name:     "word fills the line",
			words:    []string{"abc"},
			maxWidth: 3,
			want:     []string{"abc"},
		},
		{
			name:     "single short word",
			words:    []string{"a"},
			maxWidth: 3,
			want:     []string{"a  "},
		},
		{
			// "dddd" does not fit after "a b c" (that line would be 10 wide),
			// so the first line is a, b, c. Five spare spaces over two gaps
			// gives the left gap 3 and the right gap 2.
			name:     "uneven gaps",
			words:    []string{"a", "b", "c", "dddd"},
			maxWidth: 8,
			want:     []string{"a   b  c", "dddd    "},
		},
		{
			name:     "last line left justified",
			words:    []string{"a", "b", "c"},
			maxWidth: 8,
			want:     []string{"a b c   "},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := fullJustify(tt.words, tt.maxWidth)
			if !slices.Equal(got, tt.want) {
				t.Fatalf("fullJustify(%q, %d) = %q, want %q", tt.words, tt.maxWidth, got, tt.want)
			}
		})
	}
}
