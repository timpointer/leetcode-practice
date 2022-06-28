package main

import "testing"

func Test_maxDepth(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"e1", args{root: NewTree([]int{3, 9, 20, -1, -1, 15, 7})}, 3},
		{"e2", args{root: NewTree([]int{1, -1, 2})}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxDepth(tt.args.root); got != tt.want {
				t.Errorf("maxDepth() = %v, want %v", got, tt.want)
			}
		})
	}
}
