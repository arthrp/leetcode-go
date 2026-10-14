package leetcode

// 76 - Minimum Window Substring
// minWindow returns the shortest substring of s containing every character of t,
// duplicates included, or "" if none exists.
// Time is O(len(s)+len(t)). Extra memory is O(1).
func minWindow(s string, t string) string {
	var need [256]int // how many more of each byte the window needs; negative is surplus
	for i := 0; i < len(t); i++ {
		need[t[i]]++
	}
	missing := len(t) // characters of t not yet covered by the window

	bestStart, bestLen := 0, len(s)+1 // bestLen starts past any real window
	for left, right := 0, 0; right < len(s); right++ {
		if need[s[right]] > 0 { // this byte covers one still-missing character of t
			missing--
		}
		need[s[right]]--

		for missing == 0 { // s[left:right+1] covers t; shrink it from the left
			if right-left+1 < bestLen {
				bestStart, bestLen = left, right-left+1
			}
			need[s[left]]++
			if need[s[left]] > 0 { // dropping s[left] uncovers a required character
				missing++
			}
			left++
		}
	}
	if bestLen > len(s) {
		return ""
	}
	return s[bestStart : bestStart+bestLen]
}
