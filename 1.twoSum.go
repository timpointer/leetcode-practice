package main

func twoSum(nums []int, target int) []int {
	result := []int{}
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			r := nums[i] + nums[j]
			if r == target {
				result = append(result, i)
				result = append(result, j)
			}
		}
		println(i)
	}
	return result
}
