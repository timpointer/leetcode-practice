package sort

import "fmt"

var debug = true

func InsertSort(list []int) []int {
	for i := 1; i < len(list); i++ {
		temp := list[i]
		position := i - 1

		for position >= 0 {
			if list[position] > temp {
				list[position+1] = list[position]
				position--
			} else {
				break
			}
		}
		list[position+1] = temp
	}
	return list
}

func SelectionSort(list []int) []int {
	for i := 0; i < len(list)-1; i++ {
		// find lowest
		lowest := i
		for j := i + 1; j < len(list); j++ {
			if list[lowest] > list[j] {
				lowest = j
			}
		}

		if lowest != i {
			list[i], list[lowest] = list[lowest], list[i]
		}
		DPrintf("%v\n", list)
		DPrintIndex(list, i, "in")
		DPrintIndex(list, lowest, "lo")
	}
	return list
}

func BubbleSort(list []int) []int {
	unsortedUntilIndex := len(list) - 1
	sorted := false

	for sorted == false {
		sorted = true
		for i := 0; i < unsortedUntilIndex; i++ {
			if list[i] > list[i+1] {
				DPrintf("%v\n", list)
				DPrintIndex(list, unsortedUntilIndex, "un")
				DPrintIndex(list, i, "bi")
				DPrintIndex(list, i+1, "sm")
				list[i], list[i+1] = list[i+1], list[i]
				sorted = false
			}
		}
		unsortedUntilIndex--
	}

	return list
}

func DPrintf(format string, a ...any) {
	if debug {
		fmt.Printf(format, a...)
	}
}

func DPrintIndex(list []int, start int, char string) {
	if debug {
		for i, _ := range list {
			if i == 0 {
				fmt.Printf("[")
			} else {
				fmt.Printf(" ")
			}
			if i == start {
				fmt.Printf(char)
				continue
			}
			fmt.Printf("--")
		}
		fmt.Printf("]")

		fmt.Printf("\n")
	}
}
