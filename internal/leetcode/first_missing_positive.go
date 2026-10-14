package leetcode

// 41 - First Missing Positive
// firstMissingPositive returns the smallest positive integer absent from nums.
// It reorders nums in place. Time is O(len(nums)). Extra memory is O(1).
func firstMissingPositive(nums []int) int {
	n := len(nums)
	// The answer is in [1, n+1], so only values 1..n matter. Swap each such
	// value v into slot v-1 until slot i holds a value with no slot or a
	// duplicate of what its slot already has.
	for i := 0; i < n; i++ {
		for v := nums[i]; v >= 1 && v <= n && nums[v-1] != v; v = nums[i] {
			nums[i], nums[v-1] = nums[v-1], v // seats v for good, so at most n swaps overall
		}
	}
	for i, v := range nums {
		if v != i+1 { // slot i lacks its value, so i+1 is missing
			return i + 1
		}
	}
	return n + 1 // nums is a permutation of 1..n
}
