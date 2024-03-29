package main

import "testing"

func Test_romanToInt(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{"III"}, 3},
		{"2", args{"IV"}, 4},
		{"3", args{"IX"}, 9},
		{"4", args{"LVIII"}, 58},
		{"5", args{"MCMXCIV"}, 1994},

		{"6", args{"XXVII"}, 27},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := romanToInt(tt.args.s); got != tt.want {
				t.Errorf("romanToInt() = %v, want %v", got, tt.want)
			}
		})
	}
}
