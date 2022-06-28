package main

import (
	"reflect"
	"testing"
)

func TestListToList(t *testing.T) {
	type args struct {
		input []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"1", args{input: []int{1, 2, 3}}, []int{1, 2, 3}},
		{"2", args{input: []int{}}, []int{}},
		{"3", args{input: []int{1}}, []int{1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ListToList(tt.args.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ListToList() = %v, want %v", got, tt.want)
			}
		})
	}
}
