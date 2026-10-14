package leetcode

// 44 - Wildcard Matching
// isMatch reports whether the entire string s matches the wildcard pattern p.
//
// '?' matches any single character. '*' matches any sequence of characters,
// including the empty sequence. The match must cover all of s, not a substring.
//
// It scans s and p left to right with two indices, i into s and j into p.
// Literals and '?' advance both. A '*' is first assumed to match nothing:
// its index is saved in star, and the position in s where it begins is saved
// in mark. If a later character fails to match, the most recent '*' is made
// to absorb one more character of s: mark moves forward by one, i restarts
// there, and j restarts just after the '*'.
//
// Only the most recent '*' ever needs to be retried. The pattern text between
// two stars is matched at the earliest place in s where it fits, and an
// earlier place is never worse: any characters an earlier '*' could have
// absorbed can be absorbed by the later '*' instead.
//
// It uses O(1) extra memory. Time is O(len(s)+len(p)) on typical inputs and
// O(len(s)*len(p)) in the worst case, such as s = "aaaa", p = "*aab".
func isMatch(s string, p string) bool {
	i, j := 0, 0        // next unmatched character in s and in p
	star, mark := -1, 0 // index of the last '*' in p (-1 if none), and where in s it started
	for i < len(s) {
		switch {
		case j < len(p) && p[j] == '*':
			// Checked first so a '*' in p is never compared as a literal.
			// Record the star and let it match nothing for now.
			star, mark = j, i
			j++
		case j < len(p) && (p[j] == '?' || p[j] == s[i]):
			i++ // one character of s consumed by one character of p
			j++
		case star >= 0:
			// Mismatch: let the last '*' absorb one more character of s,
			// then retry the rest of the pattern right after that '*'.
			mark++
			i = mark
			j = star + 1
		default:
			return false // mismatch and no '*' to fall back on
		}
	}
	// s is fully consumed; whatever remains of p must be stars matching nothing.
	for j < len(p) && p[j] == '*' {
		j++
	}
	return j == len(p)
}
