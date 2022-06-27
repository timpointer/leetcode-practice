package sort

import (
	"fmt"
	"math/rand"
	"reflect"
	"testing"
)

func TestSort(t *testing.T) {
	//var sort = BubbleSort
	//var sort = SelectionSort
	var sort = InsertSort
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
			list: []int{65, 55, 45, 35, 25, 15, 10},
		}, []int{10, 15, 25, 35, 45, 55, 65}},
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
