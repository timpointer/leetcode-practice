package main

import "testing"

func Test_maxSubArray(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"e1", args{nums: []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}}, 6},
		{"e2", args{nums: []int{1}}, 1},
		{"e3", args{nums: []int{5, 4, -1, 7, 8}}, 23},
		{"e4", args{nums: []int{-1}}, -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxSubArray(tt.args.nums); got != tt.want {
				t.Errorf("maxSubArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
