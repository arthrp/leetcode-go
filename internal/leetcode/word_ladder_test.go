package leetcode

import "testing"

func TestLadderLength(t *testing.T) {
	tests := []struct {
		name      string
		beginWord string
		endWord   string
		wordList  []string
		want      int
	}{
		{
			name:      "shortest ladder",
			beginWord: "hit",
			endWord:   "cog",
			wordList:  []string{"hot", "dot", "dog", "lot", "log", "cog"},
			want:      5,
		},
		{
			name:      "end word missing",
			beginWord: "hit",
			endWord:   "cog",
			wordList:  []string{"hot", "dot", "dog", "lot", "log"},
			want:      0,
		},
		{
			name:      "one step",
			beginWord: "a",
			endWord:   "c",
			wordList:  []string{"c"},
			want:      2,
		},
		{
			name:      "no one-letter path",
			beginWord: "hot",
			endWord:   "dog",
			wordList:  []string{"dog"},
			want:      0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ladderLength(tt.beginWord, tt.endWord, tt.wordList)
			if got != tt.want {
				t.Fatalf("ladderLength(%q, %q, %v) = %d, want %d", tt.beginWord, tt.endWord, tt.wordList, got, tt.want)
			}
		})
	}
}
