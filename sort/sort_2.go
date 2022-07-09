package sort

func SelectionSort2(list []int) []int {
	if len(list) == 0 {
		return list
	}

	for last := len(list) - 1; last > 0; last-- {
		b := 0

		for j := 1; j <= last; j++ {
			if list[b] < list[j] {
				b = j
			}
		}
		list[b], list[last] = list[last], list[b]
	}

	return list
}

func SelectionSort3(list []int) []int {
	SelectionSort3Helper(list, len(list)-1)
	return list
}

func SelectionSort3Helper(list []int, last int) {
	if len(list) == 0 {
		return
	}

	if last <= 0 {
		return
	}
	b := 0
	for j := 1; j <= last; j++ {
		if list[b] < list[j] {
			b = j
		}
	}
	list[b], list[last] = list[last], list[b]

	SelectionSort3Helper(list, last-1)
}

func MergeSort(list []int) []int {
	return MergeSortHelper(list)
}

func MergeSortHelper(list []int) []int {
	var sl1, sl2 []int
	if len(list)-0 > 1 {
		mid := len(list) / 2
		sl1 = MergeSortHelper(list[:mid])
		sl2 = MergeSortHelper(list[mid:])
		res := merge(sl1, sl2)
		return res
	} else {
		return list
	}
}

func merge(si []int, sj []int) []int {
	DPrintf("%v %v\n", si, sj)
	res := []int{}
	var i, j int
	for {
		if i == len(si) && j == len(sj) {
			break
		}

		if i >= len(si) {
			res = append(res, sj[j])
			j++
			continue
		}
		if j >= len(sj) {
			res = append(res, si[i])
			i++
			continue
		}

		if si[i] < sj[j] {
			res = append(res, si[i])
			i++
		} else {
			res = append(res, sj[j])
			j++
		}
	}
	return res
}
