package leetcode

// 124 - Binary Tree Maximum Path Sum
// maxPathSum returns the largest sum of node values along any non-empty path
// in the tree. The path does not have to pass through the root.
// Time is O(n) for n nodes. Extra memory is O(h) for a tree of height h.
func maxPathSum(root *TreeNode) int {
	best := root.Val // root is never nil, and one node is a path

	// gain returns the best sum of a path that starts at node and goes down,
	// or 0 if every such path is negative, so the parent leaves it out.
	var gain func(node *TreeNode) int
	gain = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		left, right := gain(node.Left), gain(node.Right)
		// Every path has one highest node. The best one peaking here joins
		// the best leg down each side, where a leg of 0 is left out.
		best = max(best, node.Val+left+right)
		// A path through the parent can go down only one side.
		return max(0, node.Val+max(left, right))
	}

	gain(root)
	return best
}
