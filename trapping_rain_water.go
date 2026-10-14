package leetcode

// 42 - Trapping Rain Water
// trap returns how much water the elevation map height can hold after rain.
// Each bar is one unit wide. Time is O(len(height)). Extra memory is O(1).
func trap(height []int) int {
	left, right := 0, len(height)-1
	// Tallest bars seen from each end, including the pointer's own bar.
	// Water over a bar is the shorter of those two maxima minus the bar.
	leftMax, rightMax := height[left], height[right]
	water := 0 // fits in an int: n times the tallest bar is about 2e9
	for left < right {
		// The shorter side is the water level: a bar on the other side is at
		// least this tall, so it caps this side. A taller bar raises the max.
		// A shorter one holds the difference. A single bar, a flat map, and a
		// strictly rising or falling map add nothing.
		if leftMax <= rightMax {
			left++
			if height[left] > leftMax {
				leftMax = height[left]
			} else {
				water += leftMax - height[left]
			}
		} else {
			right--
			if height[right] > rightMax {
				rightMax = height[right]
			} else {
				water += rightMax - height[right]
			}
		}
	}
	return water
}
