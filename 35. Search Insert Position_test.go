package main

import "testing"

func Test_searchInsert(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"e1", args{
			nums:   []int{1, 3, 5, 6},
			target: 5,
		}, 2},
		{"e2", args{
			nums:   []int{1, 3, 5, 6},
			target: 2,
		}, 1},
		{"e1", args{
			nums:   []int{1, 3, 5, 6},
			target: 7,
		}, 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := searchInsert(tt.args.nums, tt.args.target); got != tt.want {
				t.Errorf("searchInsert() = %v, want %v", got, tt.want)
			}
		})
	}
}
