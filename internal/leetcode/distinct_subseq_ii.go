package leetcode

// 940 - Distinct Subsequences II
// distinctSubseqII returns the number of distinct non-empty subsequences of s, modulo 10^9+7.
// Time is O(len(s)). Extra memory is O(1).
func distinctSubseqII(s string) int {
	const mod = 1_000_000_007
	var end [26]int // end[c] is how many distinct subsequences currently end with letter c
	total := 0      // sum of end: every non-empty subsequence ends with exactly one letter
	for i := 0; i < len(s); i++ {
		c := s[i] - 'a'
		// Every subsequence so far, plus the empty one, can take this c on the
		// end, and all of those end with c. The 1 is the empty one, which
		// becomes the subsequence that is just this character.
		next := total + 1
		if next == mod {
			next = 0
		}
		// next replaces end[c] rather than adding to it. The old strings ending
		// in c are an earlier c appended to what existed before it, and this c
		// appended to the same strings spells them again, so next already
		// counts them. Every value is in [0, mod), so total ends up in
		// (-mod, 2*mod).
		total += next - end[c]
		if total < 0 {
			total += mod
		} else if total >= mod {
			total -= mod
		}
		end[c] = next // other letters keep their counts; those subsequences do not use this c
	}
	return total
}
