package main

import (
	"reflect"
	"testing"
)

func Test_getIntersectionNode(t *testing.T) {
	n1 := &ListNode{
		Val:  1,
		Next: nil,
	}
	n2 := &ListNode{
		Val:  1,
		Next: n1,
	}
	n3 := &ListNode{
		Val:  1,
		Next: n1,
	}
	type args struct {
		headA *ListNode
		headB *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{"e1", args{
			headA: n2,
			headB: n3,
		}, n1},
		{"e2", args{
			headA: NewLinkList([]int{1}),
			headB: NewLinkList([]int{2}),
		}, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getIntersectionNode(tt.args.headA, tt.args.headB); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getIntersectionNode() = %v, want %v", got, tt.want)
			}
		})
	}
}
