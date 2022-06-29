package main

import "testing"

func Test_hasCycle(t *testing.T) {
	n1 := &ListNode{
		Val:  1,
		Next: nil,
	}
	n2 := &ListNode{
		Val:  2,
		Next: n1,
	}
	n1.Next = n2
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"e1", args{head: NewLinkList([]int{})}, false},
		{"e2", args{head: n1}, true},
		{"e3", args{head: NewLinkList([]int{1})}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hasCycle(tt.args.head); got != tt.want {
				t.Errorf("hasCycle() = %v, want %v", got, tt.want)
			}
		})
	}
}
