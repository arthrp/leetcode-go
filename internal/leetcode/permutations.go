package leetcode

// 46 - Permutations
// permute returns every ordering of the distinct integers in nums.
//
// It searches by backtracking. At each step the current permutation is the
// prefix stored in path. used[i] is true when nums[i] is already in that
// prefix, so each value is chosen at most once. When the prefix is as long
// as nums, it is one complete permutation.
//
// path is one slice reused for every branch. A finished permutation must be
// copied before it is stored; otherwise later appends and pops would change
// rows already saved in result.
func permute(nums []int) [][]int {
	result := make([][]int, 0)
	path := make([]int, 0, len(nums))
	used := make([]bool, len(nums))

	var backtrack func()
	backtrack = func() {
		if len(path) == len(nums) {
			perm := make([]int, len(path))
			copy(perm, path)
			result = append(result, perm)
			return
		}

		for i := range nums {
			if used[i] {
				continue
			}
			used[i] = true               // mark nums[i] chosen
			path = append(path, nums[i]) // append it to the current prefix
			backtrack()                  // recurse to fill the rest
			path = path[:len(path)-1]    // undo the append
			used[i] = false              // undo the mark so the next candidate is tried from this prefix
		}
	}

	backtrack()
	return result
}
