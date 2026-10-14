package leetcode

// TreeNode is one node of a binary tree.
type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

// buildTree builds a tree from LeetCode's level-order form, where each item
// is an int or nil for a missing child. An empty vals gives a nil tree.
func buildTree(vals []any) *TreeNode {
	if len(vals) == 0 || vals[0] == nil {
		return nil
	}
	root := &TreeNode{Val: vals[0].(int)}
	queue := []*TreeNode{root}
	// Each node off the queue takes the next two items as its children. Items
	// left once the queue runs out have no parent, so they are ignored.
	for i := 1; i < len(vals) && len(queue) > 0; i += 2 {
		node := queue[0]
		queue = queue[1:]
		if vals[i] != nil {
			node.Left = &TreeNode{Val: vals[i].(int)}
			queue = append(queue, node.Left)
		}
		if i+1 < len(vals) && vals[i+1] != nil {
			node.Right = &TreeNode{Val: vals[i+1].(int)}
			queue = append(queue, node.Right)
		}
	}
	return root
}
