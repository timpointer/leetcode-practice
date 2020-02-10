package main

import (
	"fmt"
	"testing"
)

func Test_isPalindrome(t *testing.T) {
	type args struct {
		x int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"1", args{121}, true},
		{"2", args{-121}, false},
		{"3", args{10}, false},
		{"4", args{1001}, true},
	}
	for i, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fmt.Printf("case %d\n", i)
			if got := isPalindrome(tt.args.x); got != tt.want {
				t.Errorf("isPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
