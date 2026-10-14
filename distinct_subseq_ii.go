package leetcode

// 940 - Distinct Subsequences II
// distinctSubseqII returns the number of distinct non-empty subsequences of s,
// modulo 10^9+7.
//
// A subsequence keeps the relative order of the characters it uses. Two
// subsequences that spell the same string are counted once.
//
// end[c] is how many distinct subsequences currently end with letter c.
// Every non-empty subsequence ends with exactly one letter, so the answer is
// the sum of end, kept in total as end changes.
//
// For each character c, every subsequence seen so far, plus the empty one,
// can take a new c on the end. That set is 1+total, and every member ends
// with c, so it replaces end[c]. Replacing, rather than adding, drops the
// copies that already ended with an earlier c: those are the same strings as
// "append c to what existed before that earlier c". Counts for other letters
// stay as they are; those subsequences do not use this new character.
//
// Time is O(len(s)). Extra memory is O(1).
func distinctSubseqII(s string) int {
	const mod = 1_000_000_007
	var end [26]int
	total := 0
	for i := 0; i < len(s); i++ {
		c := s[i] - 'a'
		// 1 is the subsequence that is just this character.
		next := total + 1
		if next == mod {
			next = 0
		}
		// Every value is in [0, mod), so total ends up in (-mod, 2*mod).
		total += next - end[c]
		if total < 0 {
			total += mod
		} else if total >= mod {
			total -= mod
		}
		end[c] = next
	}
	return total
}
