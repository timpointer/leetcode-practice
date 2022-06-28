package main

import "fmt"

func maxSubArray(nums []int) int {
	res := -10000
	curSum := 0
	for _, num := range nums {
		curSum = max(curSum+num, num)
		res = max(res, curSum)
		fmt.Printf("c% 2d n% 2d r% 2d\n", curSum, num, res)
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
