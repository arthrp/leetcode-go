package leetcode

// 25 - Reverse Nodes in k-Group
// reverseKGroup reverses every group of k nodes and returns the new head.
//
// A dummy node sits in front of head so the first group is reversed the same
// way as every later group. From the node before a group, walk k steps. If
// that walk falls off the list, fewer than k nodes remain and they stay in
// order.
//
// Otherwise the group is reversed in place. Each node's Next is pointed at
// the node that previously preceded it. The walk stops at the node after the
// group, so the remainder stays attached: the last node of the reversed group
// already points there. After the reverse, the old kth node is the new head
// of the group and the old first node is the new tail. The node before the
// group is linked to that new head, and the next group starts from the new
// tail.
//
// Node values are never written. Time is O(n). Extra memory is O(1).
func reverseKGroup(head *ListNode, k int) *ListNode {
	dummy := &ListNode{Next: head}
	groupPrev := dummy // node immediately before the group about to be reversed

	for {
		kth := groupPrev
		for i := 0; i < k; i++ {
			kth = kth.Next
			if kth == nil {
				return dummy.Next // fewer than k nodes left; leave them as they are
			}
		}
		groupNext := kth.Next // first node after this group; stays put during the reverse

		// Point each node at the one that used to precede it. prev starts at
		// groupNext so the new tail remains linked to the rest of the list.
		prev := groupNext
		curr := groupPrev.Next
		for curr != groupNext {
			next := curr.Next
			curr.Next = prev
			prev = curr
			curr = next
		}

		// groupPrev.Next was the old head and is now the tail of this group.
		// kth is the new head. The next group begins after the new tail.
		newTail := groupPrev.Next
		groupPrev.Next = kth
		groupPrev = newTail
	}
}
