package leetcode

import (
	"reflect"
	"testing"
)

func TestBuildTree(t *testing.T) {
	tests := []struct {
		name string
		vals []any
		want *TreeNode
	}{
		{
			name: "missing children on one level",
			vals: []any{-10, 9, 20, nil, nil, 15, 7},
			want: &TreeNode{
				Val:  -10,
				Left: &TreeNode{Val: 9},
				Right: &TreeNode{
					Val:   20,
					Left:  &TreeNode{Val: 15},
					Right: &TreeNode{Val: 7},
				},
			},
		},
		{
			name: "gap moves later items to the next parent",
			vals: []any{1, nil, 2, 3},
			want: &TreeNode{
				Val:   1,
				Right: &TreeNode{Val: 2, Left: &TreeNode{Val: 3}},
			},
		},
		{name: "items with no parent are ignored", vals: []any{1, nil, nil, 2}, want: &TreeNode{Val: 1}},
		{name: "null root", vals: []any{nil}, want: nil},
		{name: "empty", vals: []any{}, want: nil},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := buildTree(tt.vals); !reflect.DeepEqual(got, tt.want) {
				t.Fatalf("buildTree(%v) does not match the expected tree", tt.vals)
			}
		})
	}
}
