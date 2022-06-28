package main

import (
	"reflect"
	"testing"
)

func Test_mergeTwoLists(t *testing.T) {
	NewLinkList([]int{1, 2, 4})
	type args struct {
		list1 *ListNode
		list2 *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{"e1", args{
			list1: NewLinkList([]int{1, 2, 4}),
			list2: NewLinkList([]int{1, 3, 4}),
		}, NewLinkList([]int{1, 1, 2, 3, 4, 4})},
		{"e2", args{
			list1: NewLinkList([]int{}),
			list2: NewLinkList([]int{}),
		}, NewLinkList([]int{})},
		{"e3", args{
			list1: NewLinkList([]int{}),
			list2: NewLinkList([]int{0}),
		}, NewLinkList([]int{0})},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mergeTwoLists(tt.args.list1, tt.args.list2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mergeTwoLists() = %v, want %v", got, tt.want)
			}
		})
	}
}
