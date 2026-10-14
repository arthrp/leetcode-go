package leetcode

// 154 - Find Minimum in Rotated Sorted Array II
// findMinII returns the smallest value in nums, a rotated sorted array that
// may hold duplicates.
// Time is O(log n) for n = len(nums) if the values are distinct, and O(n) in
// the worst case, when nearly all are equal. Extra memory is O(1).
func findMinII(nums []int) int {
	// nums[lo..hi] always holds the minimum. Like nums, it is sorted or has
	// one drop, a value smaller than the one before it, and every value after
	// the drop is at most every value before it.
	lo, hi := 0, len(nums)-1
	for lo < hi {
		mid := lo + (hi-lo)/2 // lo <= mid < hi
		switch {
		case nums[mid] < nums[hi]:
			// A drop after mid would make nums[hi] at most nums[mid]. So
			// nums[mid..hi] is sorted, and nums[mid] is its smallest value.
			hi = mid
		case nums[mid] > nums[hi]:
			// nums[mid..hi] is out of order, so the drop is after mid, and
			// the value right after the drop is the window's minimum.
			lo = mid + 1
		default:
			// Equal values hide which side the drop is on: it is before mid
			// in [3,1,3,3,3] and after mid in [3,3,3,1,3]. Either way,
			// nums[mid] matches nums[hi] and stays in the window, so leaving
			// out nums[hi] loses nothing.
			hi--
		}
	}
	return nums[lo]
}
