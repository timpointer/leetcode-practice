package main

import "testing"

func Test_isSymmetric(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"e1", args{root: NewTree([]int{1, 2, 2, 3, 4, 4, 3})}, true},
		{"e2", args{root: NewTree([]int{1, 2, 2, -1, 3, -1, 3})}, false},
		{"e3", args{root: NewTree([]int{1, 2, 2, 2, -1, 2})}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isSymmetric(tt.args.root); got != tt.want {
				t.Errorf("isSymmetric() = %v, want %v", got, tt.want)
			}
		})
	}
}
