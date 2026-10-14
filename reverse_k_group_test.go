package leetcode

import (
	"fmt"
	"slices"
	"testing"
)

func TestReverseKGroup(t *testing.T) {
	tests := []struct {
		vals []int
		k    int
		want []int
	}{
		{vals: []int{1, 2, 3, 4, 5}, k: 2, want: []int{2, 1, 4, 3, 5}},
		{vals: []int{1, 2, 3, 4, 5}, k: 3, want: []int{3, 2, 1, 4, 5}},
		{vals: []int{1, 2, 3, 4, 5}, k: 1, want: []int{1, 2, 3, 4, 5}},
		{vals: []int{1}, k: 1, want: []int{1}},
		{vals: []int{1, 2, 3, 4}, k: 4, want: []int{4, 3, 2, 1}},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v k=%d", tt.vals, tt.k), func(t *testing.T) {
			got := listVals(reverseKGroup(buildList(tt.vals), tt.k))
			if !slices.Equal(got, tt.want) {
				t.Fatalf("reverseKGroup(%v, %d) = %v, want %v", tt.vals, tt.k, got, tt.want)
			}
		})
	}
}

func buildList(vals []int) *ListNode {
	dummy := &ListNode{}
	tail := dummy
	for _, v := range vals {
		tail.Next = &ListNode{Val: v}
		tail = tail.Next
	}
	return dummy.Next
}

func listVals(head *ListNode) []int {
	var vals []int
	for head != nil {
		vals = append(vals, head.Val)
		head = head.Next
	}
	return vals
}
