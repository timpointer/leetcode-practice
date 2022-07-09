package sort

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_merge(t *testing.T) {
	type args struct {
		sl1 []int
		sl2 []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"e1", args{
			sl1: []int{1},
			sl2: []int{2},
		},
			[]int{1, 2},
		},
		{"e2", args{
			sl1: []int{2},
			sl2: []int{1},
		},
			[]int{1, 2},
		},
		{"e3", args{
			sl1: []int{1, 4},
			sl2: []int{2, 3},
		},
			[]int{1, 2, 3, 4},
		},
		{"e1", args{
			sl1: []int{1, 3},
			sl2: []int{2, 4},
		},
			[]int{1, 2, 3, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, merge(tt.args.sl1, tt.args.sl2), "merge(%v, %v)", tt.args.sl1, tt.args.sl2)
		})
	}
}
