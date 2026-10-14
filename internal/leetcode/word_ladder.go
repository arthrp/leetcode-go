package leetcode

// 127 - Word Ladder
// ladderLength returns the number of words in the shortest transformation
// from beginWord to endWord, or 0 when no such sequence exists.
func ladderLength(beginWord string, endWord string, wordList []string) int {
	// Each dictionary word is a node. Two words share an edge when they differ
	// by a single letter.
	dict := make(map[string]struct{}, len(wordList))
	for _, word := range wordList {
		dict[word] = struct{}{}
	}
	// If endWord is missing from wordList, no sequence exists.
	if _, ok := dict[endWord]; !ok {
		return 0
	}

	// beginWord need not be in wordList. Mark it visited at the start so a
	// copy of it in the dictionary is not reused. The length counts both ends,
	// so the search starts at 1.
	queue := []string{beginWord}
	visited := map[string]struct{}{beginWord: {}}
	length := 1

	// Breadth-first search finds the shortest path because every edge has cost 1.
	for len(queue) > 0 {
		nextLevel := make([]string, 0)
		for _, word := range queue {
			if word == endWord {
				return length
			}
			// Keep a candidate only when it is in the dictionary and unvisited.
			for _, neighbor := range neighbors(word) {
				if _, ok := dict[neighbor]; !ok {
					continue
				}
				if _, seen := visited[neighbor]; seen {
					continue
				}
				visited[neighbor] = struct{}{}
				nextLevel = append(nextLevel, neighbor)
			}
		}
		queue = nextLevel
		length++
	}
	return 0
}

// neighbors returns every word that differs from word by exactly one letter.
func neighbors(word string) []string {
	letters := []byte(word)
	out := make([]string, 0, len(letters)*25)
	for i := range letters {
		orig := letters[i]
		// Try every other letter in this position.
		for c := byte('a'); c <= 'z'; c++ {
			if c == orig {
				continue
			}
			letters[i] = c
			out = append(out, string(letters))
		}
		letters[i] = orig
	}
	return out
}
