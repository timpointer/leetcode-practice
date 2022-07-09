package sort

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"reflect"
	"testing"
)

func TestMid(t *testing.T) {
	assert.Equal(t, Mid(0, 2), 1, "they should be equal")
	assert.Equal(t, Mid(0, 3), 1, "they should be equal")
	assert.Equal(t, Mid(0, 4), 2, "they should be equal")
	assert.Equal(t, Mid(0, 5), 2, "they should be equal")
}
func Mid(a, b int) int {
	return (a + b) / 2
}

func TestSort(t *testing.T) {
	//var sort = BubbleSort
	//var sort = SelectionSort
	//var sort = InsertSort
	//var sort = SelectionSort2
	//var sort = SelectionSort3
	var sort = MergeSort
	debug = true
	type args struct {
		list []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"", args{
			list: []int{65, 55, 35, 45, 25, 15, 10},
		}, []int{10, 15, 25, 35, 45, 55, 65}},
		{"e1", args{
			list: []int{1},
		}, []int{1}},
		{"e2", args{
			list: []int{2, 1},
		}, []int{1, 2}},
		{"e3", args{
			list: []int{3, 2, 1},
		}, []int{1, 2, 3}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sort(tt.args.list); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BubbleSort() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBubbleSort_cal_bigO(t *testing.T) {
	list := createRandList(100, 1000)
	fmt.Printf("%v\n", list)
	debug = false
	BubbleSort(list)
}

func createRandList(number, limit int) []int {
	m := map[int]struct{}{}
	for i := 0; i < number; i++ {
		data := rand.Intn(limit)
		m[data] = struct{}{}
	}
	var list []int
	for n, _ := range m {
		list = append(list, n)
	}

	return list
}
