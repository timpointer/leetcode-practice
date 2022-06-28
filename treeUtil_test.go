package main

import (
	"reflect"
	"testing"
)

func TestTreeToTree(t *testing.T) {
	type args struct {
		list []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"e1", args{list: []int{1, 2, 3}}, []int{2, 1, 3}},
		{"e2", args{list: []int{1, 2, 3, 4, 5, 6, 7}},
			[]int{4, 2, 5, 1, 6, 3, 7}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TreeToTree(tt.args.list); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TreeToTree() = %v, want %v", got, tt.want)
			}
		})
	}
}
