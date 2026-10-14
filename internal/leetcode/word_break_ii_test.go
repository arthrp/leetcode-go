package leetcode

import (
	"slices"
	"strings"
	"testing"
)

func TestWordBreakII(t *testing.T) {
	tests := []struct {
		name     string
		s        string
		wordDict []string
		want     []string
	}{
		{
			name:     "example 1",
			s:        "catsanddog",
			wordDict: []string{"cat", "cats", "and", "sand", "dog"},
			want:     []string{"cats and dog", "cat sand dog"},
		},
		{
			name:     "example 2",
			s:        "pineapplepenapple",
			wordDict: []string{"apple", "pen", "applepen", "pine", "pineapple"},
			want:     []string{"pine apple pen apple", "pineapple pen apple", "pine applepen apple"},
		},
		{
			name:     "example 3",
			s:        "catsandog",
			wordDict: []string{"cats", "dog", "sand", "and", "cat"},
			want:     []string{},
		},
		{
			name:     "single letter",
			s:        "a",
			wordDict: []string{"a"},
			want:     []string{"a"},
		},
		{
			name:     "words reused",
			s:        "aaaa",
			wordDict: []string{"a", "aa"},
			want:     []string{"a a a a", "aa a a", "a aa a", "a a aa", "aa aa"},
		},
		{
			name:     "word longer than s",
			s:        "ab",
			wordDict: []string{"abc", "a", "b"},
			want:     []string{"a b"},
		},
		{
			name:     "no letter matches",
			s:        "xyz",
			wordDict: []string{"a"},
			want:     []string{},
		},
		{
			// Every run of a's splits, but no word covers the b.
			name:     "dead end at the last letter",
			s:        strings.Repeat("a", 19) + "b",
			wordDict: []string{"a", "aa", "aaa"},
			want:     []string{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := wordBreakII(tt.s, tt.wordDict)
			if !sameSentences(got, tt.want) {
				t.Fatalf("wordBreakII(%q, %q) = %q, want %q", tt.s, tt.wordDict, got, tt.want)
			}
		})
	}
}

func TestWordBreakIICount(t *testing.T) {
	// Splitting n a's into a's and aa's writes n as an ordered sum of 1s and
	// 2s. There are F(n+1) such sums, so 20 a's give F(21) = 10946.
	s, wordDict := strings.Repeat("a", 20), []string{"a", "aa"}
	if got := len(wordBreakII(s, wordDict)); got != 10946 {
		t.Fatalf("len(wordBreakII(%q, %q)) = %d, want 10946", s, wordDict, got)
	}
}

// sameSentences reports whether a and b hold the same sentences, ignoring order.
func sameSentences(a, b []string) bool {
	a, b = slices.Clone(a), slices.Clone(b)
	slices.Sort(a)
	slices.Sort(b)
	return slices.Equal(a, b)
}
