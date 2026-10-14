package leetcode

// 60 - Permutation Sequence
// getPermutation returns the kth (counting from 1) permutation of the digits
// 1..n in lexicographic order. Time is O(n^2). Extra memory is O(n).
func getPermutation(n int, k int) string {
	digits := []byte("123456789")[:n] // digits not yet placed, in increasing order
	fact := 1
	for i := 2; i < n; i++ {
		fact *= i
	}
	k-- // count from 0 so k / fact indexes digits directly
	out := make([]byte, 0, n)
	for i := n - 1; i >= 0; i-- {
		// fact is i!, the number of permutations that follow each choice of
		// the next digit, so rank k lies in block k / fact.
		idx := k / fact
		k %= fact // rank within that block
		out = append(out, digits[idx])
		digits = append(digits[:idx], digits[idx+1:]...) // shifting keeps the rest sorted; O(n)
		if i > 0 {
			fact /= i
		}
	}
	return string(out)
}
