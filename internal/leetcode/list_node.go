package leetcode

// ListNode is one node of a singly linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// buildList returns a list whose values are vals, in order.
func buildList(vals []int) *ListNode {
	dummy := &ListNode{}
	tail := dummy
	for _, v := range vals {
		tail.Next = &ListNode{Val: v}
		tail = tail.Next
	}
	return dummy.Next
}

// listVals returns the values of the list that starts at head.
func listVals(head *ListNode) []int {
	var vals []int
	for head != nil {
		vals = append(vals, head.Val)
		head = head.Next
	}
	return vals
}
