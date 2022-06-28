package main

import (
	"reflect"
	"testing"
)

func Test_inorderTraversal(t *testing.T) {
	e1Node := &TreeNode{
		Val:  1,
		Left: nil,
		Right: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 3,
			},
		},
	}
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"e1", args{root: e1Node}, []int{1, 3, 2}},
		{"e2", args{root: nil}, []int{}},
		{"e3", args{root: &TreeNode{
			Val:   1,
			Left:  nil,
			Right: nil,
		}}, []int{1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := inorderTraversal(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("inorderTraversal() = %v, want %v", got, tt.want)
			}
		})
	}
}
