package sort

import (
	"reflect"
	"testing"
)

func TestReversal(t *testing.T) {
	type args struct {
		txt string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"", args{txt: "abcde"}, "edcba"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Reversal(tt.args.txt); got != tt.want {
				t.Errorf("Reversal() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCountX(t *testing.T) {
	type args struct {
		txt string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"", args{txt: "axaxax"}, 3},
		{"", args{txt: "axaxa"}, 2},
		{"", args{txt: "aba"}, 0},
		{"", args{txt: ""}, 0},
		{"", args{txt: "xaax"}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CountX(tt.args.txt); got != tt.want {
				t.Errorf("CountX() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCountStep(t *testing.T) {
	type args struct {
		stair int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"0", args{stair: 0}, 1},
		{"1", args{stair: 1}, 1},
		{"2", args{stair: 2}, 2},
		{"3", args{stair: 3}, 4},
		{"4", args{stair: 4}, 7},
		{"5", args{stair: 5}, 13},
		{"6", args{stair: 6}, 24},
		{"7", args{stair: 20}, 24},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CountStep2(tt.args.stair); got != tt.want {
				t.Errorf("CountStep() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAnagram_banchmark(t *testing.T) {
	println(len(Anagram("abcdefghik")))
}

func TestAnagram(t *testing.T) {
	type args struct {
		txt string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{"name", args{txt: "abc"}, []string{
			"abc",
			"acb",
			"bac",
			"bca",
			"cab",
			"cba",
		}},

		{"name", args{txt: ""}, []string{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Anagram(tt.args.txt); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Anagram() = %v, want %v", got, tt.want)
			}
		})
	}
}
