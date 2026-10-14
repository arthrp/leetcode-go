package leetcode

import "strings"

// 140 - Word Break II
// wordBreakII returns every way to add spaces to s so that each word is in
// wordDict.
// Time is O(m + n*L^2*2^n) for n = len(s), L the longest word length, and m
// the total length of wordDict. Extra memory is O(len(wordDict) + n), plus the
// sentences returned.
func wordBreakII(s string, wordDict []string) []string {
	dict := make(map[string]struct{}, len(wordDict))
	maxLen := 0 // no word is longer, so s[i:j] can be a word only when j-i <= maxLen
	for _, word := range wordDict {
		dict[word] = struct{}{}
		maxLen = max(maxLen, len(word))
	}
	// isWord reports whether s[i:j] is a dictionary word.
	isWord := func(i, j int) bool {
		_, found := dict[s[i:j]]
		return found
	}

	n := len(s)
	// ok[i] reports whether s[i:] splits into dictionary words. It does when
	// some word s[i:j] is followed by a suffix s[j:] that splits, so the table
	// fills from the right, starting with the empty suffix, which splits into
	// no words.
	ok := make([]bool, n+1)
	ok[n] = true
	for i := n - 1; i >= 0; i-- {
		for j := i + 1; j <= min(n, i+maxLen); j++ {
			if ok[j] && isWord(i, j) {
				ok[i] = true
				break
			}
		}
	}

	result := make([]string, 0)   // sentences may appear in any order
	words := make([]string, 0, n) // the sentence so far; it spells s[:start] when search(start) runs
	var search func(start int)
	search = func(start int) {
		if start == n { // the words spell all of s
			result = append(result, strings.Join(words, " "))
			return
		}
		for end := start + 1; end <= min(n, start+maxLen); end++ {
			// Take s[start:end] only when it is a word and the rest of s after
			// it splits too. Then the search never follows a dead end, and
			// every word it takes leads to at least one sentence.
			if !ok[end] || !isWord(start, end) {
				continue
			}
			words = append(words, s[start:end])
			search(end)
			words = words[:len(words)-1] // drop the word and try a longer one
		}
	}
	search(0)
	return result
}
